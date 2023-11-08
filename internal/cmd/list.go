package cmd

import "github.com/armnerd/zcache/internal/expire"

// Lpush 将一个值插入到列表头部
func Lpush(key string, value string, extra ...string) {
	// 看是否容器中已经存在
	// 检查是否过期
	// 过期需重建
	for k := range extra {
		if k == EXTRA_EXPIRE {
			expire.Record(key, extra[k], expire.LIST)
		}
	}
}

// Lpop 移出并获取列表的第一个元素
func Lpop(key string) (member string) {
	// 检查是否过期
	return
}

// Rpush 将一个值插入到列表尾部
func Rpush(key string, value string, extra ...string) {
	// 看是否容器中已经存在
	// 检查是否过期
	// 过期需重建
	for k := range extra {
		if k == EXTRA_EXPIRE {
			expire.Record(key, extra[k], expire.LIST)
		}
	}
}

// Rpop 移出并获取列表的最后一个元素
func Rpop(key string) (member string) {
	// 检查是否过期
	return
}

// Lrange 获取列表指定范围内的元素
func Lrange(key string, start string, stop string) string {
	// 检查是否过期
	return "ok"
}

// Llen 获取列表长度
func Llen(key string) string {
	// 检查是否过期
	return "ok"
}
