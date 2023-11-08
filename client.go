package zcache

import (
	"fmt"
	"io"
	"net"

	"github.com/aceld/zinx/znet"
)

type cmd struct {
	first  string
	second string
	third  string
	fourth string
}

var input cmd

type Client struct {
	opts *Options
}

func NewClient(opts ...OptionFunc) *Client {
	options := loadOptions(opts...)
	c := &Client{
		opts: options,
	}
	return c
}

func (c *Client) Run() {
	fmt.Println("Client start")
	url := fmt.Sprintf("%s:%d", c.opts.Addr, c.opts.Port)
	conn, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		// wait
		fmt.Print("> : ")
		fmt.Scanln(&input.first, &input.second, &input.third, &input.fourth)
		command := fmt.Sprintf("%s %s %s %s", input.first, input.second, input.third, input.fourth)
		// send and echo
		echo := c.sendMsg(conn, command)
		fmt.Printf("$ : %s\n", echo)

		// clean buffer
		input = cmd{}
	}
}

func (c *Client) sendMsg(conn net.Conn, command string) string {
	// 发封包 message 消息
	dp := znet.NewDataPack()
	msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte(command)))
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("write error err ", err)
		return ""
	}

	// 先读出流中的 head 部分
	headData := make([]byte, dp.GetHeadLen())
	_, err = io.ReadFull(conn, headData) // ReadFull 会把 msg 填充满为止
	if err != nil {
		fmt.Println("read head error")
		return ""
	}
	// 将 headData 字节流 拆包到 msg 中
	msgHead, err := dp.Unpack(headData)
	if err != nil {
		fmt.Println("server unpack err:", err)
		return ""
	}

	if msgHead.GetDataLen() > 0 {
		// msg 是有 data 数据的，需要再次读取 data 数据
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		// 根据 dataLen 从 io 中读取字节流
		_, err := io.ReadFull(conn, msg.Data)
		if err != nil {
			fmt.Println("server unpack data err:", err)
			return ""
		}

		var ret = string(msg.Data)
		return ret[1 : len(ret)-1]
	}
	return ""
}
