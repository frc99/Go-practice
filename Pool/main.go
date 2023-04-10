package main

import (
	"fmt"
	"runtime"
	"time"
)

//var wg sync.WaitGroup

type Job interface {
	DoSomething()
}

type worker struct {
	jobChan chan Job
	//IsIdle  bool
}

func NewWorker() worker {
	return worker{
		jobChan: make(chan Job),
		//IsIdle:  true,
	}
}
func (w *worker) Run(wq chan chan Job) {
	go func() {
		//将worker通道注册到
		wq <- w.jobChan
		for {
			select {
			case job := <-w.jobChan:
				job.DoSomething()
			}
		}

		/*for job := range w.jobChan {
			job.DoSomething()
		}*/
	}()
}
func (w *worker) Report() {

}

type workerPool struct {
	workerNum   int
	jobQueue    chan Job //接收外部的任务
	workerQueue chan chan Job
}

func NewWorkerPool(workNum int) *workerPool {
	return &workerPool{
		workerNum:   workNum,
		jobQueue:    make(chan Job),
		workerQueue: make(chan chan Job, workNum),
	}
}
func (wp *workerPool) Run() {
	//初始化worker
	for i := 0; i < wp.workerNum; i++ {
		w := NewWorker()
		w.Run(wp.workerQueue)
	}
	//schedule():循环获取可用的worker,往worker中写job
	go func() {
		for {
			select {
			case job := <-wp.jobQueue:
				idleworker := <-wp.workerQueue
				idleworker <- job
			}

		}
	}()

}

// 实现job接口
type Code struct {
	Num int
}

func (c *Code) DoSomething() {
	fmt.Println("code num:", c.Num)
	time.Sleep(1 * time.Second)
}

func main() {
	workerNum := 100 * 100 * 2
	p := NewWorkerPool(workerNum)
	p.Run()

	//写数据
	dataNum := 100 * 100 * 10
	go func() {
		for i := 0; i < dataNum; i++ {
			c := &Code{Num: i}
			p.jobQueue <- c
		}
		close(p.jobQueue)
	}()
	//wg.Wait()
	//close(p.workerQueue)

	//维持程序运行
	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(5 * time.Second)
	}

}
