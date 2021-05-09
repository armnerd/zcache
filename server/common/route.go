package common

import (
	"errors"
	"reflect"
	"strings"
)

// config
type config struct {
	handlerMap map[string]interface{}
}

// Setting 全局配置
var Setting config

// Init 初始化配置
func Init() {
	Setting.handlerMap = make(map[string]interface{})

	// string
	Setting.handlerMap["set"] = Set
	Setting.handlerMap["get"] = Get
	Setting.handlerMap["del"] = Del

	// hash
	Setting.handlerMap["hset"] = Hset
	Setting.handlerMap["hget"] = Hget
	Setting.handlerMap["hgetall"] = Hgetall
	Setting.handlerMap["hkeys"] = Hkeys
	Setting.handlerMap["hvals"] = Hvals
	Setting.handlerMap["hdel"] = Hdel

	// list
	Setting.handlerMap["lpush"] = Lpush
	Setting.handlerMap["lpop"] = Lpop
	Setting.handlerMap["rpush"] = Rpush
	Setting.handlerMap["rpop"] = Rpop
	Setting.handlerMap["lrange"] = Lrange
	Setting.handlerMap["llen"] = Llen

	// set
	Setting.handlerMap["sadd"] = Sadd
	Setting.handlerMap["smembers"] = Smembers
	Setting.handlerMap["spop"] = Spop
	Setting.handlerMap["srem"] = Srem

	// zset
	Setting.handlerMap["zadd"] = Zadd
	Setting.handlerMap["zrangebyscore"] = Zrangebyscore
	Setting.handlerMap["zscore"] = Zscore
	Setting.handlerMap["zrem"] = Zrem
}

// Handler 功能分发
func Handler(args []string) interface{} {
	// deal params and call function
	var target = strings.ToLower(args[0])
	var found bool
	_, found = Setting.handlerMap[target]
	if !found {
		return "[unknown command]"
	}
	args = shift(args)
	ret, err := call(Setting.handlerMap, target, args)
	if err != nil {
		return err.Error()
	}
	return ret
}

// 函数反射
func call(m map[string]interface{}, name string, params []string) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("[params number wrong]")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}

// 删除数组开头元素
func shift(arr []string) []string {
	var res []string
	for index, one := range arr {
		if index != 0 {
			res = append(res, one)
		}
	}
	return res
}
