package note

import (
	"fmt"
)

func increase(n *int) {
	*n++ // n = n + 1
	fmt.Printf("\nincrease结束时n=%v\nn的内存地址为%v\nn指向的值为%v\n", n, &n, *n)
}

// 2.4 指针
func Pointer() {
	var src = 2022
	increase(&src)
	fmt.Printf("\n调用increase(ptr)之后，src=%v\nsrc的内存地址为%v\n", src, &src)
	var ptr = new(int)
	fmt.Printf("\nptr=%v\nptr的内存地址为%v\nptr指向的值为%v\n", ptr, &ptr, *ptr)
}

// 2.5 fmt格式化
func FmtVerbs() {
	fmt.Println("\n2.5.1 通用")
	fmt.Printf("%%\n")

	fmt.Println("\n2.5.2 整数")
	i := 123
	fmt.Printf("%U\n", i)
	fmt.Printf("%c\n", i)
	fmt.Printf("%q\n", i)

	fmt.Println("\n2.5.3 浮点数")
	f := 123.456
	fmt.Printf("%f\n", f)
	fmt.Printf("%.2f\n", f)
	fmt.Printf("%20f\n", f)
	fmt.Printf("%b\n", f)
	fmt.Printf("%e\n", f)
	fmt.Printf("%E\n", f)
	fmt.Printf("%G\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%X\n", f)

	fmt.Println("\n2.5.4 布尔")
	fmt.Printf("%t\n", f == 123.456)

	fmt.Println("\n2.5.5 字符串或byte切片")
	s := "hello world"
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)

	fmt.Println("\n2.5.6指针")
	p := &s
	fmt.Printf("%p\n", p)
}

// 2.6 运算符
func Operator() {
	fmt.Println("\n2.6.1 算数运算符")
	fmt.Printf("8%%3=%d\n", 8%3)
	i := 123
	i++ // i = i + 1
	fmt.Printf("i=%d\n", i)

	fmt.Println("\n2.6.2 位运算符")
	var b uint8 = 0b00111100
	fmt.Printf("b>>2=%b\n", b>>2)
	fmt.Printf("b<<2=%b\n", b<<2)
	var b1 uint8 = 0b00111100
	var b2 uint8 = 0b11001111
	fmt.Printf("b1&b2=%b\n", b1&b2)
	fmt.Printf("b1|b2=%b\n", b1|b2)
	fmt.Printf("b1^b2=%b\n", b1^b2)

	fmt.Println("\n2.6.3 赋值运算符")
	b += 3 // b = b + 3
	fmt.Printf("b=%d\n", b)

	fmt.Println("\n2.6.4 关系运算符")
	fmt.Printf("b1==b2?%t\n", b1 == b2)

	fmt.Println("\n逻辑运算符")
	bool1 := true
	bool2 := false
	fmt.Printf("bool1&&bool2?%t\n", bool1 && bool2)
	fmt.Printf("boll1 || bool2?%t\n", bool1 || bool2)
	fmt.Printf("!bool2?%t\n", !bool2)
}

// 3.1 if...else
func IfElse() {
	var age uint8
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("小朋友不要学编程")
	} else if age < 25 {
		fmt.Println("大朋友不要学编程")
	} else {
		fmt.Println("老朋友不要学编程")
	}
}

// 3.2 switch case
func SwitchCase() {
	var weekday uint8
	fmt.Println("请输入星期（数字）")
	fmt.Scanln(&weekday)
	switch weekday {
	case 1:
		fmt.Println("小朋友不要学编程")
	case 2:
		fmt.Println("大朋友不要学编程")
	default:
		fmt.Println("老朋友不要学编程")
	}

}

// 3.3 for循环
func For() {
	fmt.Println("\n3.3.1 无线循环")
	i := 1
	for {
		fmt.Print(i, "\t")
		i++
		if i == 11 {
			fmt.Println()
			break
		}
	}
	fmt.Println("\n3.3.2 条件循环")
	i = 1
	for i < 11 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println()
	fmt.Println("\n3.3.3 标准for循环")
	for j := 1; j < 11; j++ {
		fmt.Print(j, "\t")
	}
	fmt.Println()
}

// 3.4 label 与 goto
func LabelAndGoto() {
	fmt.Println("\n 3.4.1 label")
outside:
	for i := 0; i < 10; i++ {
		for j := 0; j <= 10; j++ {
			fmt.Print("+ ")
			if j == 4 && i == 9 {
				break outside
			}
		}
		fmt.Println()
	}
	fmt.Println("\n 3.4.2 goto")
	fmt.Print("1 ")
	if i := 1; i == 1 {
		goto four
	}
	fmt.Print("2 ")
	fmt.Print("3 ")
four:
	fmt.Print("4 ")
	fmt.Print("5 ")
	fmt.Print("6 ")

}

// 3.5 函数
// func getRes(n1, n2 int) (sum, difference int) {
// 	sum = n1 + n2
// 	difference = n1 - n2
// 	return
// }

func Function() {
	res1, res2 := func(n1, n2 int) (sum, difference int) {
		sum = n1 + n2
		difference = n1 - n2
		return
	}(2, 3)
	fmt.Println("res1=", res1, ", res2=", res2)

}

// func Function() {
// 	var getRes = func (n1, n2 int) (sum, difference int) {
// 		sum = n1 + n2
// 		difference = n1 - n2
// 		return
// 	}
// 	res1, res2 := getRes(2, 3)
// 	fmt.Println("res1=", res1, ", res2=", res2)

// 	// fmt.Printf("getRes=%v,Type of getRes=%T\n", getRes, getRes)
// }

// 3.6 defer
func deferUtil() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函数被第%v次调用\n", i)
		return i
	}
}
func Defer() int {
	f := deferUtil()
	defer f(1)
	defer f(2)
	defer f(3)
	return f(4)
}

func DeferRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

func Array() {
	// 4.1.1 声明
	// var a [3]int = [3]int{
	// 	1,
	// 	345,
	// 	789,
	// }
	var a = [...]int{
		123,
		456,
		789,
	}
	a[0] = 123
	fmt.Println("\nfor 遍历")
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%v]=%v\n", i, a[i])
	}
	fmt.Println("\nfor range 遍历")
	for i, v := range a {
		fmt.Printf("a[%v]=%v\n", i, v)
	}
	fmt.Println("\n4.1.3 多维数组")
	var twoDimensionalArray [3][4]int = [3][4]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
	}
	for i, v := range twoDimensionalArray {
		for i2, v2 := range v {
			fmt.Printf("a[%v][%v]=%v\t", i, i2, v2)
		}
		fmt.Println()
	}
}

// 4.2 切片
func Slice() {
	array := [5]int{1, 2, 3, 4, 5}
	var s1 []int = array[1:4] //  [开始引用的index:结束引用的index+1)//[0:len(arrary)]等效于[:]
	s1[0] = 0
	fmt.Println("array=", array)
	s2 := s1[1:]
	s2[0] = 0
	fmt.Println("array=", array)
	var s3 []int
	fmt.Println("does s3 == nil?", s3 == nil)
	s3 = make([]int, 3, 5) // make ([]Type, len, cap)
	fmt.Printf("len(s3)=%v,cap(s3)=%v\n", len(s3), cap(s3))
	s4 := []int{1, 2, 3} // 由系统自动创建底层数组
	fmt.Println("s4=", s4)

	fmt.Println("\n4.2.3 追加元素")
	s1 = append(s1, 6, 7, 8) // 底层创建了新的数组，不在引用原数组。
	s1[4] = 0
	fmt.Println("array=", array)
	fmt.Println("s1=", s1)
	s5 := append(s1, s2...)
	fmt.Println("s5=", s5)

	fmt.Println("\n4.2.4, 复制数组")
	s6 := []int{1, 1}
	copy(s5, s6) // 容量能接收多少，就接收多少。
	fmt.Println("s5", s5)

	fmt.Println("\n4.2.5 string与[]byte")
	str := "hello 世界"
	fmt.Printf("[]byte(str)=%v\n[]byte(str)=%s\n", []byte(str), []byte(str))
	for i, v := range str {
		fmt.Printf("str[%d]=%c\n", i, v)
	}
}

// 4.3 map
func Map() {
	var m1 map[string]string
	fmt.Println("m1 ==nil ?", m1 == nil)
	m1 = make(map[string]string) // make(Type,初始size)
	m1["早上"] = "敲代码"
	m1["中午"] = "送外卖"
	m1["晚上"] = "开滴滴"
	fmt.Println("m1 =", m1)
}

//
