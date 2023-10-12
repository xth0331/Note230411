package main

import "fmt"

func main() {
	// 定义一个channel
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine 运行...")
		c <- 666 // 将666发送给c
	}()
	num := <-c //从c中接受数据，赋值给num
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")

}
