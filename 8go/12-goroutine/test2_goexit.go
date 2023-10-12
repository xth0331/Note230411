package main

import (
	"fmt"
	"time"
)

func main() {
	// 用go创建承载一个形参为空
	/* 	go func() {
		defer fmt.Println("A.defer")
		// return
		func() {
			defer fmt.Println("B.defer")
			// 推出当前goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}() */
	go func(a int, b int) bool {
		fmt.Println("a= ", a, "b=", b)
		return true

	}(10, 20)
	//
	for {
		time.Sleep(1 * time.Second)
	}
}
