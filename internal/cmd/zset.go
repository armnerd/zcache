package cmd

import (
	"strconv"

	"github.com/armnerd/zcache/internal/data"
	"github.com/armnerd/zcache/pkg/zset"
)

// Zadd 向有序集合添加一个成员，或者更新已存在成员的分数
func Zadd(key string, score string, member string) string {
	scoreInt, _ := strconv.Atoi(score)
	res, found := data.ZsetContainer[key]
	if !found {
		res = zset.NewZset()
		res.Put(member, scoreInt)
		data.ZsetContainer[key] = res
	} else {
		res.Put(member, scoreInt)
	}
	return "ok"
}

// Zrangebyscore 通过分数返回有序集合指定区间内的成员
func Zrangebyscore(key string, min string, max string) string {
	minInt, _ := strconv.Atoi(min)
	maxInt, _ := strconv.Atoi(max)
	res, found := data.ZsetContainer[key]
	if !found {
		return "not found"
	}
	return res.Members(minInt, maxInt)
}

// Zscore 返回有序集中，成员的分数值
func Zscore(key string, member string) string {
	res, found := data.ZsetContainer[key]
	if !found {
		return "not found"
	}
	return res.GetScore(member)
}

// Zrem 移除有序集合中的一个成员
func Zrem(key string, member string) string {
	res, found := data.ZsetContainer[key]
	if !found {
		return "not found"
	}
	return res.RemoveMember(member)
}
