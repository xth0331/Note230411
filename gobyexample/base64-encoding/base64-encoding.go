package main

import (
	b64 "encoding/base64"
	"fmt"
)

// Go 提供了对Base64编解码的内建支持。
// 这些语法引入了encoding/base64包，并使用别名b64代替默认的base64。

func main() {
	// 要编解码的字符串。
	data := "abc123!?$*&()'-=@~"
	// Go 同时支持标准 base64 以及 URL 兼容 base64。 这是使用标准编码器进行编码的方法。
	// 编码器需要一个 []byte，因此我们将 string 转换为该类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码可能会返回错误，如果不确定输入信息格式是否正确， 那么，你就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()
	// 使用URL base64格式进行编解码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
