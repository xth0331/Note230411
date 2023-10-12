package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// close可以关闭一个channel
		close(c)

	}()
	// for {
	// 	// ok如果为true表示channel没有关闭，如果false表示channel关闭
	// 		if data, ok := <-c; ok {
	// 			fmt.Println(data)
	// 		} else {
	// 			break
	// 		}

	// }
	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("Main finished ..")
}
