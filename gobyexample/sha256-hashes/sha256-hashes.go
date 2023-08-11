package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256散列（hash）经常用于生成二进制文件或文件块的短标识。例如，TLS/SSL证书使用SHA256来计算一个证书的签名。

// Go在多个crypot/*包中实现了一系列散列函数。
func main() {
	s := "sha256 this string"
	// 这里我们从一个新的散列开始。
	h := sha256.New()
	//要写入处理的字节。如果时一个字符串，需要使用[]bypte(s)强制转换成字节数组。
	h.Write([]byte(s))
	// Sum得到最终的散列值字符串切片。Sum接收一个参数，可以用来给现有的字符串切片追加额外的字节切片：但是一般都不需要这样做。
	bs := h.Sum(nil)
	// SHA256值经常以16进制格式输出。例如在git commit中。我们这里也使用%x来将散列格式化为16进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
