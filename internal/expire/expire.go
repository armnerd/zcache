package expire

import (
	"strconv"
	"time"
)

type ContainerType string

const (
	STRING ContainerType = "string"
	HASH   ContainerType = "hash"
	LIST   ContainerType = "list"
	SET    ContainerType = "set"
	ZSET   ContainerType = "zset"
)

type ExpireInfo struct {
	Key       string        // 储存 key
	Expire    int64         // 过期时间
	Container ContainerType // 储存类型，在哪个容器里
	Index     int           // 在 TimeMachine 中的下标
}

var (
	TimeMachine  []ExpireInfo
	KeyExpireMap map[string]ExpireInfo
)

func init() {
	KeyExpireMap = make(map[string]ExpireInfo)
}

// 记录过期时间
func Record(key string, expire string, container ContainerType) {
	e, err := strconv.ParseInt(expire, 10, 64)
	if err != nil {
		return
	}
	now := time.Now().Unix()
	record := ExpireInfo{
		Key:       key,
		Expire:    now + e,
		Container: container,
		Index:     len(TimeMachine),
	}
	TimeMachine = append(TimeMachine, record)
	KeyExpireMap[key] = record
}

// 定时清理
func Clean(gap int) {

}
