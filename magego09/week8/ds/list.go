package main

import (
	"container/list"
	"fmt"
)

// 6 4 2 1 3 5
func testList() {
	l := list.New() // 双向链表
	l.PushBack(1)
	l.PushBack(3)
	l.PushBack(5)
	l.PushFront(2)
	l.PushFront(4)
	l.PushFront(6)

	tail := l.Back() // 尾部元素
	fmt.Println(tail.Value)
	for {
		tail = tail.Prev()
		if tail == nil {
			break
		}
		fmt.Println(tail.Value)
	}
	fmt.Println("===========================")
	head := l.Front() // 首元素
	fmt.Println(head.Value)
	for {
		head = head.Next()
		if head == nil {
			break
		}
		fmt.Println(head.Value)
	}

}
