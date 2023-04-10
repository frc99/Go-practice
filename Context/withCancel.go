package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go Speak(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	//time.Sleep(1 * time.Second)
}

func Speak(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	for range ticker.C {
		//for {
		fmt.Println("hello", i)
		i++
		select {
		case <-ctx.Done():
			//fmt.Println("--------------")
			fmt.Println("我要闭嘴了")
			return
		default:
			fmt.Println("balabalabalabala")
		}
	}
}

func Shout(ctx context.Context) {

}
