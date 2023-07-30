package main

import "fmt"

func main() {
	// 这里make了一个字符串通道，最多允许缓存2个值。
	messages := make(chan string, 2)
	// 由于次通道是有缓冲的，因此我们可以将这些值发送到通道中，而不需并发的接收。
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
