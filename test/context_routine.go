package main

import (
	"context"
	"jyutil"

	//"sync"
	"time"
)

//範例: 兄弟執行緒間不求同生只求同死，使用context

const shortDuration = 1001 * time.Millisecond

//var wg sync.WaitGroup //宣告計數器

func aRoutine(ctx context.Context) {
	defer wg.Done()

	select {
	case <-time.After(1 * time.Second):
		jyutil.Output("aRoutine - Success!")
	case <-ctx.Done():
		jyutil.Output("aRoutine - ", ctx.Err()) // context deadline exceeded
	}
}

func testContext() {
	timeout := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), timeout) // 宣告一個context.Context並設定1.001秒之內未執行完的執行緒將發出ctx.Err
	defer cancel()
	go aRoutine(ctx) // 啟動aRoutine執行序
	wg.Add(1)        // 計數器+1
	wg.Wait()        // 等待計數器歸零
}
