package main

import (
	"fmt"
	"os"
)

// 命令行参数是指定程序运行参数的一个常见方式。例如，go run hello.go，程序go使用了run和hello.go两个参数。
func main() {
	// os.Args提供原始命令行参数访问功能。注意，切片的第一个参数是该程序的路径，而os.Args[1:]保存了程序全部的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	// 可以使用标准的下标取得单个参数的值。
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
