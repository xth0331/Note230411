package main

import (
	"fmt"
)

/*
声明函数：

	func 函数名称() {
		功能代码
	}
*/
// func printLing(n int) {
// 	for i := 1; i <= n; i++ {
// 		for k := 1; k <= n-i; k++ {
// 			fmt.Print(" ")
// 		}

// 		for j := 1; j <= 2*i-1; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}

// 	for i := n - 1; i >= 1; i-- {
// 		for k := 1; k <= n-i; k++ {
// 			fmt.Print(" ")
// 		}

// 		for j := 1; j <= 2*i-1; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }
// func leiJia(x, y int) {
// 	var ret = 0
// 	for i := x; i <= y; i++ {
// 		ret += i
// 	}
// 	fmt.Println(ret)
// }

// //	func add(a, b int) {
// //		fmt.Print(a + b)
// //	}
// //
// //	func add2(a, b int) int{
// //		c := a + b
// //		return c
// //	}
// //
// // 无返回值的函数不能赋值给变量
// func add(s ...int) int {
// 	var sum = 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	return sum
// }

//	func login(user, pwd string) (bool, string) {
//		if user == "root" && pwd == "123" {
//			// 登录成功
//			return true, user
//		} else {
//			return false, ""
//		}
//	}
//
//	func main() {
//		// 函数调用
//		printLing(7)
//		leiJia(56, 1020)
//		add(9, 10, 10, 1)
//		ret := add(1, 2, 3)
//		fmt.Println(ret)
//		b, user := login("root", "123")
//		if b {
//			fmt.Println(user)
//		}
//	}
func main() {
	// 固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])

	}
	for index, value := range myArray2 {
		fmt.Println("index =", index, ",value=", value)
	}
}
