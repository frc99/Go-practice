package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	Key := "trace_id"
	Val := "2023"
	var wg sync.WaitGroup
	wg.Add(1)
	ctx := context.WithValue(context.Background(), Key, Val)
	go func(ctx context.Context) {
		defer wg.Done()

		fmt.Println("trace_id")
		fmt.Printf("%s\n", ctx.Value(Key).(string))
	}(ctx)
	wg.Wait()
	fmt.Println("over")
}
