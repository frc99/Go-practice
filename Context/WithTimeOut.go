package main

import (
	"context"
	"fmt"
	"time"
)

// 达到超时时间终止接下来的执行
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	deal(ctx, cancel)
}
func deal(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d\n", i)
			//cancel()
		}
	}
}
