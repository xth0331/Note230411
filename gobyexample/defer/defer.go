package main

import (
	"fmt"
	"os"
)

/*
Defer 用于确保程序在执行完成后，会调用某个函数，一般是执行清理工作。
Defer的用途跟其他语言的ensure或finally类似
*/
func main() {
	// 假设我们想要创建一个文件、然后写入数据、最后在程序结束时关闭该文件。
	// 这里展示了如何通过defer来做到这一切。
	//
	// 在createFile后立即得到一个文件对象，我们使用defer通过closeFile来关闭这个文件。
	// 这回在封闭函数main结束时执行，即writeFile完成后。
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
func createFile(p string) *os.File {
	fmt.Println("createing")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

// 关闭文件时，进行错误检查是非常重要的， 即使在 defer 函数中也是如此。
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
