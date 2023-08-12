package main

import (
	"bufio"
	"fmt"
	"os"
)

// 在Go中，写文件与我们在前面看过的读文件方法类似

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	// 开始！这里展示如何写入一个字符串（或者只是一些字节）到一个文件。
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)
	//对于更细粒度的写入，先打开一个文件。
	f, err := os.Create("/tmp/dat2")
	check(err)
	//打开文件后，一个喜欢行的操作是，立即使用defer调用文件的Close。
	defer f.Close()
	// 可以按期望那样Write字节切片。
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)
	// WriteString也是可用的。
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	//调用Sync将缓冲区的数据写入磁盘。
	f.Sync()
	// 与我们前面看到的带缓冲区的Reader一样，bufio还提供了带缓冲的Writer。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	//使用Flush来确保，已将所有缓冲区操作应用于底层writer。
	w.Flush()
}
