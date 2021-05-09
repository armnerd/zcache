package common

// Sadd 向集合添加一个成员
func Sadd(key string, member string) string {
	return "ok"
}

// Smembers 返回集合中的所有成员
func Smembers(key string) string {
	return "ok"
}

// Spop 移除并返回集合中的一个随机元素
func Spop(key string) string {
	return "ok"
}

// Srem 移除集合中一个成员
func Srem(key string, member string) string {
	return "ok"
}
