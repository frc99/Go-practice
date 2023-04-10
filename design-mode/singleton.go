package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	instance *Singleton
	mu       sync.Mutex
)

//var instance *Singleton  = &Singleton{}

type Singleton struct {
}

// lazy
func GetLazyInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		fmt.Println("instance...")
	})
	return instance
}

// hunger 饿汉模式将在包加载的时候就会创建单例对象
func GetHungerInstance() *Singleton {
	return instance
}

// double lock check
func GetDLInstance() *Singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}

func main() {
	var s *Singleton
	s = GetLazyInstance()
	fmt.Println(s)
	s = GetLazyInstance()
	fmt.Println(s)
}
