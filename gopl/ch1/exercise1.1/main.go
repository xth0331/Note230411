package main

import (
	"fmt"
	"os"
)

func main() {

	s, sep := "", ""
	fmt.Println(len(os.Args))
	if len(os.Args) == 1 {
		fmt.Println("没有指定命令行参数")
	} else {
		for index, arg := range os.Args[1:] {
			s += sep + arg
			sep = ""
			fmt.Printf("第%v个参数为%v，索引为%v\n", index+1, arg, index)
		}
		fmt.Printf("参数列表：%v\n", os.Args[1:])
	}
	fmt.Printf("被执行的命令为（不包含参数）%v\n", os.Args[0])

}
