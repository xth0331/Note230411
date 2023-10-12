package main

import "fmt"

func main() {
	// var a [3]int //  三个整数的数组
	// fmt.Println(a[0]) // 输出第一个元素
	// fmt.Println(a[len(a)-1])

	// 打印索引和元素
	// for i, v := range a {
	// 	fmt.Printf("%d %d\n", i, v)
	// }
	// 只打印元素
	// for _, v := range a {
	// 	fmt.Printf("%d\n", v)
	// }
	// var q [3]int = [3]int{1, 2, 3}
	q := [...]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])
	fmt.Printf("%T,%v\n", q,q) // "[3]int"
}
