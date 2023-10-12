package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p person) String() string {
	return fmt.Sprintf("Name is %s,Age is %d", p.Name, p.Age)
}
func main() {
	p := person{Name: "阿水淀粉", Age: 20}
	//fmt.Println(p.String())
	//pt := reflect.TypeOf(p)
	//// 遍历person字段
	//for i := 0; i < pt.NumField(); i++ {
	//	fmt.Println("字段", pt.Field(i).Name)
	//}
	//// 遍历person的方法
	//for i := 0; i < pt.NumMethod(); i++ {
	//	fmt.Println("方法", pt.Method(i).Name)
	//}
	//stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	//writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()
	//fmt.Println("是否实现了fmt.Stringer:", pt.Implements(stringerType))
	//fmt.Println("是否实现了io.Writer:", pt.Implements(writerType))
	// struct to json
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	// json to struct
	respJSON := "{\"name\":\"李四\",\"age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p)
}
