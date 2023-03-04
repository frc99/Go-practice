package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("10.128.180.137:%d\n", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s关闭", address)
				return
			}
			conn.Close()
			fmt.Printf("%s打开", address)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d sceonds", elapsed)
}
