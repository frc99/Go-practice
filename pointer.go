package main

import (
	"fmt"
	"sync"
)

type Connection struct {
	data     []byte
	dataType int
}

func main() {
	var ch chan *Connection
	var wg sync.WaitGroup
	ch = make(chan *Connection, 10)
	conn := &Connection{
		data:     []byte("package"),
		dataType: 1,
	}
	ch <- conn
	fmt.Println("------")
	wg.Add(1)
	go proc(ch, &wg)
	close(ch)

	//close了才可以wait否则会一直等下去
	wg.Wait()

}
func proc(ch chan *Connection, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("读取数据：", v)
	}
	if _, ok := <-ch; !ok {
		return
	}

}
