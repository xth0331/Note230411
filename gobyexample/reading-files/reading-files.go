package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读写文件在很多程序都是必须的基本任务。首先我们来看一些读文件的例子。

// 读取文件需要经常进行错误检查，这个帮方法可以精简下面的错误检查过程。
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	// 最基本的文件读取任务或许就是将文件内容读取到内存中。
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))
	// 您通常会希望对文件的读取方式和内容进行更多的控制。对于这个任务，首先使用Open打开一个文件，以获取一个os.File值。
	f, err := os.Open("/tmp/dat")
	check(err)
	// 从文件的开始位置读取一些字节。最多允许读取5个字节，但是还要注意实际读取了多少个。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	// 也可以Seek到一个文件中已知的位置，并从这个位置开始读取。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))
	// 例如，io包提供了一个更健壮的实现ReadAtLeast，用于读取上面那种文件。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s \n", n3, o3, string(b3))
	// 没有内奸的倒带，但是Seek(0,0)实现了这个功能
	_, err = f.Seek(0, 0)
	check(err)

	// bufio包实现了一个缓冲读取器，这可能有助于提高许多小读操作的效率，以及它提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	// 任务解释后要关闭这个文件（通常这个操作应该在Open操作后使用defer来完成。）

	f.Close()
}
