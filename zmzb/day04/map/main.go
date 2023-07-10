package main

import "fmt"

func main() {
	// var stus map[string]string
	// fmt.Println(stus, reflect.TypeOf(stus))
	// 声明并初始化
	// var x = []int{1, 2, 3} // 切片的声明并初始化
	var stus = map[string]string{"name": "yuan", "age": "32"}
	// 支持key查询 map对象[key]
	fmt.Println(stus)
	fmt.Println(stus["name"])
	fmt.Println(stus["age"])

	//写入一个key-value
	stus["gender"] = "male"
	fmt.Println(stus) // map[age:32 gender:male name:yuan]
	stus["height"] = "180cm"
	// 修改一个key-value
	stus["height"] = "190cm"

	delete(stus, "gender")
	fmt.Println(stus)

	// var s []int
	s := make([]int, 3)
	s[0] = 100
	// var stu02 map[string]string // map 引用类型
	var stu02 = make(map[string]string)
	stu02["name"] = "rain"
	stu02["age"] = "30"
	stu02["gender"] = "male"

	fmt.Println(stu02)
	for k, v := range stu02 {
		fmt.Println(k, v)
	}
}
