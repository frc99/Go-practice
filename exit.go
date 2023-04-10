package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done的值:", <-done)
				return
			default:
				fmt.Println("监控中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	close(done)
	fmt.Println("程序退出")
}
