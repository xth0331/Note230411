package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go 的 math/rand 包提供了伪随机数生成器。

func main() {
	// 例如，rand.Intn返回一个随机的整数n，且0<=n<100.
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	// rand.Float64返回一个64位浮点数f，且0.0<=f<1.0.
	fmt.Println(rand.Float64())
	// 这个技巧可以用来生成其他范围的随机浮点数，例如5.0<=f<10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()
	// 默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列。 要产生不同的数字序列，需要给定一个不同的种子。
	// 注意，对于想要加密的随机数，使用此方法并不安全， 应该使用 crypto/rand。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// 调用上面返回的rand.Rand，就像调用rand包中的函数一样。
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	// 如果使用相同的种子生成随机数生成器，会生成相同的随机数序列。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
