package common

import (
	"encoding/json"
	"fmt"

	"github.com/armnerd/zcache/server/hash"
)

// Hset 将哈希表 key 中的字段 field 的值设为 value
func Hset(key string, field string, value string) string {
	res, found := HashContainer[key]
	if !found {
		res = hash.New()
		res.Put(field, value)
		HashContainer[key] = res
	} else {
		res.Put(field, value)
	}
	return "ok"
}

// Hget 获取存储在哈希表中指定字段的值
func Hget(key string, field string) string {
	res, found := HashContainer[key]
	if !found {
		return "not found"
	}
	var value, err = res.Get(field)
	if !err {
		return "not found"
	}
	return fmt.Sprint(value)
}

// Hgetall 获取在哈希表中指定 key 的所有字段和值
func Hgetall(key string) string {
	res, found := HashContainer[key]
	if !found {
		return "not found"
	}
	var mapInstances []map[string]interface{}
	for k, v := range res.All() {
		instance := map[string]interface{}{fmt.Sprint(k): v}
		mapInstances = append(mapInstances, instance)
	}
	jsonStr, _ := json.Marshal(mapInstances)
	return string(jsonStr)
}

// Hkeys 获取所有哈希表中的字段
func Hkeys(key string) string {
	res, found := HashContainer[key]
	if !found {
		return "not found"
	}
	jsonStr, _ := json.Marshal(res.Keys())
	return string(jsonStr)
}

// Hvals 获取哈希表中所有值
func Hvals(key string) string {
	res, found := HashContainer[key]
	if !found {
		return "not found"
	}
	jsonStr, _ := json.Marshal(res.Values())
	return string(jsonStr)
}

// Hdel 删除一个哈希表字段
func Hdel(key string, field string) string {
	res, found := HashContainer[key]
	if !found {
		return "not found"
	}
	res.Remove(field)
	return "ok"
}
