package common

// Lpush 将一个值插入到列表头部
func Lpush(key string, value string) string {
	return "ok"
}

// Lpop 移出并获取列表的第一个元素
func Lpop(key string) string {
	return "ok"
}

// Rpush 将一个值插入到列表尾部
func Rpush(key string, value string) string {
	return "ok"
}

// Rpop 移出并获取列表的最后一个元素
func Rpop(key string) string {
	return "ok"
}

// Lrange 获取列表指定范围内的元素
func Lrange(key string, start string, stop string) string {
	return "ok"
}

// Llen 获取列表长度
func Llen(key string) string {
	return "ok"
}
