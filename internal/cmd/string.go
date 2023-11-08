package cmd

import (
	"github.com/armnerd/zcache/internal/data"
	"github.com/armnerd/zcache/internal/expire"
)

// Set 设置指定 key 的值
func Set(key string, value string, extra ...string) {
	// 检查是否过期
	data.StringContainer[key] = value
	for k := range extra {
		if k == EXTRA_EXPIRE {
			expire.Record(key, extra[k], expire.STRING)
		}
	}
}

// Get 获取指定 key 的值
func Get(key string) string {
	res, found := data.StringContainer[key]
	if !found {
		return "not found!"
	}
	// 检查是否过期
	return res
}

// Del 该命令用于在 key 存在时删除 key
func Del(key string) {
	_, found := data.StringContainer[key]
	if !found {
		return
	}
	delete(data.StringContainer, key)
}
