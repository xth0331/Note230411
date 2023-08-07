package main

import "fmt"

/*
Go 通过使用 recover 内置函数，可以从 panic 中 恢复recover 。
recover 可以阻止 panic 中止程序，并让它继续执行。
*/

// 这是一个panic函数

func mayPanic() {
	panic("a problem")
}

// 必须在defer函数中调用recover。当跳出引发panic的函数时，defer会被激活，其中的recover会捕获panic。

func main() {
	// recover的返回值时在调用panic时抛出的错误。
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	// 这行代码不会执行，因为 mayPanic 函数会调用 panic。 main 程序的执行在 panic 点停止，并在继续处理完 defer 后结束。
	fmt.Println("After mayPanic()")
}
