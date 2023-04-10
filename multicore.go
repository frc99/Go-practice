package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(2)

	var w sync.WaitGroup
	count := int32(0)
	w.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				//count++
				atomic.AddInt32(&count, 1)
			}
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println(count)
}
