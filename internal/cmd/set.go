package cmd

import "github.com/armnerd/zcache/internal/expire"

// Sadd 向集合添加一个成员
func Sadd(key string, member string, extra ...string) {
	for k := range extra {
		if k == EXTRA_EXPIRE {
			expire.Record(key, extra[k], expire.SET)
		}
	}
}

// Smembers 返回集合中的所有成员
func Smembers(key string) (members []string) {
	return
}

// Spop 移除并返回集合中的一个随机元素
func Spop(key string) (member string) {
	return
}

// Srem 移除集合中一个成员
func Srem(key string, member string) {
	// undo
}
