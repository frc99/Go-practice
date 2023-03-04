package main

import (
	"encoding/json"
	"fmt"
)

type personInfo struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email" xml:"email"`
}

type personInfo1 struct {
	Name  string `json:"name"`
	Email string `json:"email" xml:"email"`
	C     string
}

func main() {
	// 创建数据
	p := personInfo{Name: "Piao", Age: 10, Email: "piaoyunsoft@163.com"}

	// 序列化
	data, _ := json.Marshal(&p)
	fmt.Println(string(data))

	// 反序列化
	var p1 personInfo1
	err := json.Unmarshal([]byte(data), &p1) // 貌似这种解析方法需要提前知道 json 结构
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Printf("name=%s, c=%s, email=%s\n", p1.Name, p1.C, p1.Email)
	}
	fmt.Printf("%+v\n", p1)

	// 反序列化
	//res, err := simplejson.NewJson([]byte(data))
	//if err != nil {
	//fmt.Println("err: ", err)
	//} else {
	//fmt.Printf("%+v\n", res)
	//}

	//解析未知的json
	str := `{"name":"test","product_id":"1","number":"110011","price":"0.01","is_on_sale":"true"}`
	var px interface{}
	json.Unmarshal([]byte(str), &p)
	// 现在我们需要从这个interface{}解析出里面的数据
	m := px.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Printf("%s is string, value: %s\n", k, vv)
		case int:
			fmt.Printf("%s is int, value: %d\n", k, vv)
		case int64:
			fmt.Printf("%s is int64, value: %d\n", k, vv)
		case bool:
			fmt.Printf("%s is bool, vaule: %v", k, vv)
		default:
			fmt.Printf("%s is unknow type\n", k)
		}
	}

}
