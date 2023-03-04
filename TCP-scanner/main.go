package main

import (
	"fmt"
	"net"
	"sort"
)

// wg *sync.WaitGroup
func worker(ports chan int, res chan int) {
	for p := range ports {
		//fmt.Println(p)
		address := fmt.Sprintf("10.128.180.137:%d\n", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s ERROR\n", address)
			res <- 0
			continue
		}
		//wg.Done()
		res <- p
		conn.Close()
		fmt.Printf("%s打开", address)

	}

}

func main() {
	//var wg sync.WaitGroup
	ports := make(chan int, 100) //100 worker可以从ports中取
	result := make(chan int)
	var open []int
	//ar close []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}

	go func() {
		for i := 0; i < 65535; i++ { //1024个端口
			//wg.Add(1)
			ports <- i
		}
	}()

	for i := 0; i < 65535; i++ { //1024个端口
		//这一步阻塞了，只有线程执行完才能够进行下一步，因此不需要wg
		port := <-result
		if port != 0 {
			open = append(open, port)
		}
	}
	close(ports)
	close(result)
	sort.Ints(open)
	for _, port := range open {
		fmt.Printf("%d\n", port)
	}
	//wg.Wait()
}
