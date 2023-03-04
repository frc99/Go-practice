package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go GoSleep(i, c)
	}
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopher := <-c:
			fmt.Println("GopherId:%d", gopher)
		case <-timeout:
			fmt.Println("停等两秒没等到")
			return
		}
	}
	return
}

func GoSleep(id int, c chan int) {
	//time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	time.Sleep(time.Duration(1 * time.Second))
	c <- id
}
