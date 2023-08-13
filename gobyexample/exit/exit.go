package main

import (
	"fmt"
	"os"
)

// 使用os.Exit可以立即给定的状态退出程序。
func main() {
	defer fmt.Println("!")
	os.Exit(3)
}
