package main

import (
	"fmt"
	"time"
)

/*
我们可以使用通道来同步协程之间的执行状态。
这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成。
如果需要等待多个协程，WaitGroup 是一个更好的选择。
*/

// 我们将要在协程中运行这个函数。 done 通道将被用于通知其他协程这个函数已经完成工作。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Print("done")
	// 发送一个值来通知我们已经完工啦。

	done <- true
}

func main() {
	// 运行一个worker 协程，并给予用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)
	// 程序将一直阻塞，直到收到worker使用通道发送的通知。
	<-done
}
