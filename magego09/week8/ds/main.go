package main

import "fmt"

func main() {
	//testList()
	lst := NewList()
	lst.PushBack(1)
	lst.PushBack(3)
	lst.PushBack(5)

	lst.PushFront(2)
	lst.PushFront(4)
	lst.PushFront(6)
	lst.Remove(3)
	lst.Remove(5)
	lst.Remove(6)
	lst.Remove(8)

	// 遍历
	head := lst.Head
	fmt.Println(head.Value)
	for {
		head = head.Next
		if head == nil {
			break
		}
		fmt.Println(head.Value)
	}
}
