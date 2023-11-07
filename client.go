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
	Addr string
	Port string
}

func NewClient() *Client {
	c := &Client{}
	return c
}

func (c *Client) Run() {
	fmt.Println("Client start")
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		// listen
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
	//发封包message消息
	dp := znet.NewDataPack()
	msg, _ := dp.Pack(znet.NewMsgPackage(0, []byte(command)))
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println("write error err ", err)
		return ""
	}

	//先读出流中的head部分
	headData := make([]byte, dp.GetHeadLen())
	_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
	if err != nil {
		fmt.Println("read head error")
		return ""
	}
	//将headData字节流 拆包到msg中
	msgHead, err := dp.Unpack(headData)
	if err != nil {
		fmt.Println("server unpack err:", err)
		return ""
	}

	if msgHead.GetDataLen() > 0 {
		//msg 是有data数据的，需要再次读取data数据
		msg := msgHead.(*znet.Message)
		msg.Data = make([]byte, msg.GetDataLen())

		//根据dataLen从io中读取字节流
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
