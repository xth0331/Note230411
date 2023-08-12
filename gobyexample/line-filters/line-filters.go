package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 行过滤器（line filter）是一种常见的程序类型，它读取stdin上的输入，对其进行处理，然后将处理结果打印到stdout。grep和sed就是常见的行过滤器。
// 这里是一个使用GO编写的行过滤器示例，它将所有的输入文字转化为大写的版本，你可以使用这个模式来写一个你自己的Go行过滤器。
func main() {
	// 用带缓冲的scanner包装无缓冲的os.Stdin，这为我们提供了一种方便的Scan方法，将前进到下一个令牌（默认为：下一行）。
	scanner := bufio.NewScanner(os.Stdin)

	// Text返回当前的token，这里指的是输入的下一行。
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		// 输出转换为大写后的行。
		fmt.Println(ucl)
	}
	// 检查Scan的错误，文件结束符（EOF）是可以接受的，它不会被Scan当作一个错误。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)

	}
}
