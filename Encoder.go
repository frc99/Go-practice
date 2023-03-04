package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Skill string
}

func main() {
	stu := Student{"tom", 12, "football"}
	data, err := json.Marshal(&stu)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
		return
	}
	fmt.Println("序列化后: ", string(data))
}
