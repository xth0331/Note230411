package main

import "fmt"

// 通道(channels) 是连接多个协程的管道。 你可以从一个协程将值发送到通道，然后在另一个协程中接收。
func main() {
	// 使用make(chan val-type) 创建一个新的通道。通道类型就是他们需要传递的类型。
	messages := make(chan string)
	// 使用channel <- 语法发送一个新的值到通道中。这里我们在一个新的协程中发送“ping”到上面创建的messages通道中。
	go func() { messages <- "ping" }()
	// 使用<-channel 语法从通道中接收一个值。这里我们会收到在上面发送的“ping”消息并将其打印。
	msg := <-messages
	fmt.Println(msg)
}
