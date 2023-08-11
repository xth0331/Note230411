package main

import (
	"fmt"
	"strconv"
)

//从字符串中解析数字在很多程序时一个基础常见的任务，而在go中，时这样处理的。
// 内建strconv包提供了数字解析能力。

func main() {
	// 使用ParseFloat，这里的64表示解析的数的位数。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// 在使用ParseInt解析整型数时，例子中的参考0表示自动推断字符串所表示的数字的进制。64表示返回的整型数时以64位存储的。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	// ParseInt会自动识别出字符串是十六进制数。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	// ParseUint也是可用的。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	// Atoi是一个基础的十进制整数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	// 在输入错误时，解析函数会返回一个错误。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
