package cmd

import "github.com/armnerd/zcache/internal/expire"

// Sadd 向集合添加一个成员
func Sadd(key string, member string, extra ...string) {
	// 检查是否过期
	// 过期需重建
	for k := range extra {
		if k == EXTRA_EXPIRE {
			expire.Record(key, extra[k], expire.SET)
		}
	}
}

// Smembers 返回集合中的所有成员
func Smembers(key string) (members string) {
	// 检查是否过期
	return
}

// Spop 移除并返回集合中的一个随机元素
func Spop(key string) (member string) {
	// 检查是否过期
	return
}

// Srem 移除集合中一个成员
func Srem(key string, member string) {
	// undo
}
