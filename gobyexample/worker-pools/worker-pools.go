package main

import (
	"fmt"
	"time"
)

/*
这是 worker 程序，我们会并发的运行多个 worker。
worker 将在 jobs 频道上接收工作，并在 results 上发送相应的结果。
每个 worker 我们都会 sleep 一秒钟，以模拟一项昂贵的（耗时一秒钟的）任务。
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
