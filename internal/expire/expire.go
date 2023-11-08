package expire

type ExpireInfo struct {
	Key       string // 储存 key
	Expire    int64  // 过期时间
	Container string // 储存类型，在哪个容器里
}

var (
	TimeMachine  []ExpireInfo
	KeyExpireMap map[string]ExpireInfo
)

func init() {
	KeyExpireMap = make(map[string]ExpireInfo)
}

// 记录过期时间
func Record(key string, expire int64, container string) {
	record := ExpireInfo{
		Key:       key,
		Expire:    expire,
		Container: container,
	}
	TimeMachine = append(TimeMachine, record)
	KeyExpireMap[key] = record
}

// 定时清理
func Clean(gap int) {

}
