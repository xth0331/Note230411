package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// tty: pair<type: *os.File,value: "/dev/tty"文件描述符>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	var r io.Reader
	// r : pair<type:*os.File, value: "/dev/tty"文件描述符>
	r = tty
	// w : pair<type: *os.File, value: "/dev/tty"文件描述符>
	var w io.Writer
	// w: pair<type:*os.File, value:"/dev/tty"文件描述符>
	w = r.(io.Writer)
	w.Write([]byte("HELLO THIS IS A TEST!!!\n"))
}
