package main

import "fmt"

// interface{}是万能类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
	// interface{} 如何区分引用的底层数据类型到底是什么
	// 给interface{} 提供 "断言"的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type ")

	} else {
		fmt.Println("arg is string type,value = ", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(1.23)

}
