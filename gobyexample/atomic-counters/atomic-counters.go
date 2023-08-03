package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Go 中最主要的状态管理机制是依靠通道间的通信来完成的。
我们已经在工作池的例子中看到过，并且，还有一些其他的方法来管理状态。
这里，我们来看看如何使用 sync/atomic 包在多个协程中进行 _原子计数_。
*/
func main() {
	// 使用一个无符号整型变量来表示计数器
	var ops uint64
	// WaitGroup帮助我们等待所有协程完成它们的工作。
	var wg sync.WaitGroup
	// 我们会启动50个协程，并且每个协程会嫁给你计数器递增1000次。
	for i := 0; i < 50; i++ {
		wg.Add(1)
		// 使用AddUint64来让计数器自动增加，使用&语法给定ops的内存地址。
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	// 等待，所有协程完成。
	wg.Wait()
	// 现在可以安全的访问ops，因为我们知道，此时没有协程写入ops，
	// 此外，还可以使用atomic.LoadUint64之类的函数，在原子更新的同时安全地读取它们。
	fmt.Println("ops:", ops)
}
