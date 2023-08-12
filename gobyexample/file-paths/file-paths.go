package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// filepath包为文件路径，提供了方便的跨操作系统的解析和构建函数；比如：Linux的dir/file和windows下的dir\file。
func main() {
	//应该使用Join来构建可移植的路径。它接收任意数量的参数，并参照传入顺序构造一个对应层次结构的路径。
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)
	//应该总是使用Join代替手动拼接/和\。除了可移植性，Join还会删除多余的分隔符和目录，使得路径更加规范。
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	// Dir和Base可以被用于分割路径中的目录和文件。此外，Split可以一次性调用返回上面两个函数的结果。
	fmt.Println("dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	// 判断路径是否为绝对路径。
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	filename := "config.json"
	// 某些文件名包含了扩展名（文件类型）。我们可以用Ext将扩展名分割出来。
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	// 想获取文件名清楚扩展名后的值，请使用strings.TrmSuffix。
	fmt.Println(strings.TrimSuffix(filename, ext))
	// Rel选招basepath与targetpath之间的相对路径。如果相对路径不存在，则返回错误。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

}
