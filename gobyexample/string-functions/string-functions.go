package main

//标准库的 strings 包提供了很多有用的字符串相关的函数。

import (
	"fmt"
	s "strings"
)

// 我们给fmt.Println一个较短的别名。
var p = fmt.Println

func main() {
	// 这是一些strings中有用的函数例子。由于它们都是包的函数，而不是字符串对象自身的方法，
	// 这意味着我们需要在调用函数时，将字符串作为第一个参数进行传递。
	p("Contains: ", s.Contains("test", "es"))
	p("Count:    ", s.Count("test", "t"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))
	p("Index:    ", s.Index("test", "e"))
	p("Join:     ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:   ", s.Repeat("a", 5))
	p("Replace:  ", s.Replace("foo", "o", "0", -1))
	p("Replace:  ", s.Replace("foo", "o", "0", 1))
	p("Split:    ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:  ", s.ToLower("TEST"))
	p("ToUpper:  ", s.ToUpper("test"))
	p()
	// 虽然不是strings的函数，但仍然值的一提的时，获取字符串长度以及通过索引获取一个字节的机制。
	// 上面的len以及索引工作在字节级别上。Go使用UTF-8编码字符串，因此通常按原样使用。
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

}
