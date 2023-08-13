package main

import (
	"fmt"
	"os"
	"strings"
)

// 环境变量是一种向Unix程序传递配置信息的常见方式。让我们来看看如何设置、获取以及列出环境变量。
func main() {
	// 使用os.Setenv来设置一个键值对。使用os.Getenv获取一个键对应的值。如果键不存在，将会返回一个空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	// 使用os.Environ来列出所有环境变量键值对。这个函数会返回一个KEY=value形式的字符串切片。你可以使用strings.SplitN来的到键和值。这里我们打印所有的键。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
