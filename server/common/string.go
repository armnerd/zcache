package common

// Set 设置指定 key 的值
func Set(key string, value string) string {
	StringContainer[key] = value
	return "ok"
}

// Get 获取指定 key 的值
func Get(key string) string {
	res, found := StringContainer[key]
	if !found {
		return "not found!"
	}
	return res
}

// Del 该命令用于在 key 存在时删除 key
func Del(key string) string {
	_, found := StringContainer[key]
	if !found {
		return "not found!"
	}
	delete(StringContainer, key)
	return "ok"
}
