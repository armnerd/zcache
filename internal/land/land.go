package land

import (
	"context"
	"fmt"
	"time"
)

// 定时持久化
func Land(ctx context.Context, gap time.Duration) {
	fmt.Println("开始定时持久化~")
	count := 0
	go func() {
		timeTicker := time.NewTicker(gap)
		for {
			select {
			case <-timeTicker.C:
				fmt.Printf("第%d波定时持久化~\n", count)
				count++
			case <-ctx.Done():
				fmt.Println("let's call it a day")
				return
			}
		}
	}()
}
