package route

import (
	"context"
	"fmt"
	"strings"

	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/armnerd/zcache/internal/cmd"
)

// 一线天
var commandQuene chan []string

func init() {
	commandQuene = make(chan []string, 5)
}

// Router 路由
type Router struct {
	znet.BaseRouter
}

// Handle 处理器
func (rt *Router) Handle(request ziface.IRequest) {
	var command = string(request.GetData())
	// 读取客户端的数据
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", command)

	// 处理参数, 分发处理
	var args []string
	var temp = strings.Split(command, " ")
	for _, arg := range temp {
		if arg != "" {
			args = append(args, arg)
		}
	}
	var res = Handler(args)

	// 回写消息
	var echo = fmt.Sprint(res)
	err := request.GetConnection().SendBuffMsg(0, []byte(echo))
	if err != nil {
		fmt.Println(err)
	}
}

// Handler 功能分发
func Handler(args []string) (ret interface{}) {
	if len(args) == 0 {
		ret = "no command found"
		return
	}
	first := cmd.RedisCmd(args[0])
	_, ok := cmd.CmdAvailable[first]
	if !ok {
		ret = "illegal command"
		return
	}
	_, kind := cmd.CmdWrite[first]
	if kind {
		// 写操作的异步处理
		commandQuene <- args
		ret = "done"
		return
	} else {
		// 读取类直接操作
		switch first {
		case cmd.GET:
			// undo
		case cmd.HGET:
			// undo
		case cmd.HGETALL:
			// undo
		case cmd.HKEYS:
			// undo
		case cmd.HVALS:
			// undo
		case cmd.LRANGE:
			// undo
		case cmd.LLEN:
			// undo
		case cmd.SMEMBERS:
			// undo
		case cmd.ZRANGEBYSCORE:
			// undo
		case cmd.ZSCORE:
			// undo
		}
	}
	return
}

// 顺序消费，避免冲突，和 redis 一样的套路
func OneLineSky(ctx context.Context) {
	fmt.Println("开始处理写操作~")
	go func() {
		for {
			select {
			case arg := <-commandQuene:
				first := cmd.RedisCmd(arg[0])
				switch first {
				case cmd.SET:
					// undo
				case cmd.DEL:
					// undo
				case cmd.HSET:
					// undo
				case cmd.HDEL:
					// undo
				case cmd.LPUSH:
					// undo
				case cmd.LPOP:
					// undo
				case cmd.RPUSH:
					// undo
				case cmd.RPOP:
					// undo
				case cmd.SADD:
					// undo
				case cmd.SPOP:
					// undo
				case cmd.SREM:
					// undo
				case cmd.ZADD:
					// undo
				case cmd.ZREM:
					// undo
				}
			case <-ctx.Done():
				fmt.Println("let's call it a day")
				return
			}
		}
	}()
}
