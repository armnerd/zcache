package expire

import (
	"context"
	"fmt"
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
	oldRecord, found := KeyExpireMap[key]
	if found {
		// 更新过期时间
		oldRecord.Expire = now + e
		TimeMachine[oldRecord.Index] = oldRecord
		KeyExpireMap[key] = oldRecord
	} else {
		record := ExpireInfo{
			Key:       key,
			Expire:    now + e,
			Container: container,
			Index:     len(TimeMachine),
		}
		TimeMachine = append(TimeMachine, record)
		KeyExpireMap[key] = record
	}
}

// 检查是否过时
func Check(key string) (res bool) {
	// 如果过期则删除相关所有信息
	return
}

// 定时清理
func Clean(ctx context.Context, gap time.Duration) {
	fmt.Println("定时清理开始工作~")
	count := 0
	go func() {
		timeTicker := time.NewTicker(gap)
		for {
			select {
			case <-timeTicker.C:
				fmt.Printf("第%d波定时清理~\n", count)
				count++
			case <-ctx.Done():
				fmt.Println("let's call it a day")
				return
			}
		}
	}()
}
