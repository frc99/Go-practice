package main

import (
	"log"
)

const (
	Debug = 1
)

func Dprintf(format string, v ...interface{}) {
	if Debug == 1 {
		log.Printf(format, v...)
	}
}
func CPrintf[T any](format string, v ...T) {
	if Debug == 1 {
		log.Printf(format, v)
	}
}
func main() {
	num := 1
	//num2 := 1024
	var name string
	name = "str"
	Dprintf("num:%v,name:%s", num, name)
	//CPrintf("num:%v", num, num2)
	//CPrintf("name:%v", name)
}
