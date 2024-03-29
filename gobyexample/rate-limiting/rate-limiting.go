package main

import (
	"fmt"
	"time"
)

func main() {
	//首先，我们将看一个基本的速率限制。
	//假设我们想限制对收到请求的处理，我们可以通过一个渠道处理这些请求。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// limiter通道每200ms接收一个值。这是我们任务速率限制的调度器
	limiter := time.Tick(200 * time.Millisecond)
	// 通过在每次请求前阻塞limiter通道的一次接收，可以将频率限制为，每200ms执行一次。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	// 有时候我们可能希望在速率限制方案中允许短暂的并发请求，并同时保留总体速率限制。
	// 我们可以通过缓冲通道来完成此任务。
	// burstyLimiter 通道允许最多 3 个爆发（bursts）事件。
	burstyLimiter := make(chan time.Time, 3)
	// 填充通道，表示允许的爆发（bursts）
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 每 200ms 我们将尝试添加一个新的值到 burstyLimiter中， 直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	// 现在，模仿另外5个传入请求。受益于burstyLimit的爆发能力，前三个请求可以快速完成。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
