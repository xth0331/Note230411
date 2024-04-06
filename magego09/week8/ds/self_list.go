package main

type ListNode struct {
	Value int
	Prev  *ListNode
	Next  *ListNode
}
type MyDoubleList struct {
	Head *ListNode
	Tail *ListNode
	Len  int
}

func NewList() *MyDoubleList {
	return &MyDoubleList{
		Len:  0,
		Head: nil,
		Tail: nil,
	}
}

func (lst *MyDoubleList) PushBack(v int) {
	node := &ListNode{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}
	if lst.Tail != nil {
		lst.Tail.Next = node
		node.Prev = lst.Tail
		lst.Tail = node
	} else {
		lst.Tail = node
		lst.Head = node
	}
	lst.Len += 1
}
func (lst *MyDoubleList) PushFront(v int) {
	node := &ListNode{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}
	if lst.Head != nil {
		lst.Head.Prev = node
		node.Next = lst.Head
		lst.Head = node
	} else {
		lst.Head = node
		lst.Tail = node
	}
	lst.Len += 1
}

func (lst *MyDoubleList) Remove(v int) {
	head := lst.Head
	for head != nil {
		if head.Value == v {
			lst.Len -= 1
			prev := head.Prev
			next := head.Next

			head.Prev = nil
			head.Next = nil
			if prev != nil {
				prev.Next = next
			} else {
				lst.Head = next
			}
			if next != nil {
				next.Prev = prev
			} else {
				lst.Tail = prev
			}
		}
		head = head.Next

	}
}
