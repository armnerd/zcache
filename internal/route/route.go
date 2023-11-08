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
var commandQuene chan bullet

func init() {
	commandQuene = make(chan bullet, 5)
}

type bullet struct {
	Args []string
	Res  chan string
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
	res := Handler(args)

	// 回写消息
	var echo = fmt.Sprint(res)
	err := request.GetConnection().SendBuffMsg(0, []byte(echo))
	if err != nil {
		fmt.Println(err)
	}
}

// Handler 功能分发
func Handler(args []string) (ret string) {
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
	resQ := make(chan string)
	b := bullet{
		Args: args,
		Res:  resQ,
	}
	commandQuene <- b
	ret = <-resQ
	return
}

// 顺序消费，避免冲突，和 redis 一样的套路
func OneLineSky(ctx context.Context) {
	fmt.Println("一线天开始工作~")
	go func() {
		for {
			select {
			case bullet := <-commandQuene:
				args := bullet.Args
				ret := "ok"
				first := cmd.RedisCmd(args[0])
				switch first {
				case cmd.SET:
					cmd.Set(args[1], args[2])
				case cmd.GET:
					ret = cmd.Get(args[1])
				case cmd.DEL:
					cmd.Del(args[1])
				case cmd.HSET:
					cmd.Hset(args[1], args[2], args[3])
				case cmd.HGET:
					ret = cmd.Hget(args[1], args[2])
				case cmd.HGETALL:
					ret = cmd.Hgetall(args[1])
				case cmd.HKEYS:
					ret = cmd.Hkeys(args[1])
				case cmd.HVALS:
					ret = cmd.Hvals(args[1])
				case cmd.HDEL:
					cmd.Hdel(args[1], args[2])
				case cmd.LPUSH:
					cmd.Lpush(args[1], args[2], args[3])
				case cmd.LPOP:
					ret = cmd.Lpop(args[1])
				case cmd.RPUSH:
					cmd.Rpush(args[1], args[2], args[3])
				case cmd.RPOP:
					ret = cmd.Rpop(args[1])
				case cmd.LRANGE:
					ret = cmd.Lrange(args[1], args[2], args[3])
				case cmd.LLEN:
					ret = cmd.Llen(args[1])
				case cmd.SADD:
					cmd.Sadd(args[1], args[2])
				case cmd.SPOP:
					ret = cmd.Spop(args[1])
				case cmd.SMEMBERS:
					ret = cmd.Smembers(args[1])
				case cmd.SREM:
					cmd.Srem(args[1], args[2])
				case cmd.ZADD:
					cmd.Zadd(args[1], args[2], args[3], args[4])
				case cmd.ZRANGEBYSCORE:
					ret = cmd.Zrangebyscore(args[1], args[2], args[3])
				case cmd.ZSCORE:
					ret = cmd.Zscore(args[1], args[2])
				case cmd.ZREM:
					cmd.Zrem(args[1], args[2])
				}
				bullet.Res <- ret
			case <-ctx.Done():
				fmt.Println("let's call it a day")
				return
			}
		}
	}()
}
