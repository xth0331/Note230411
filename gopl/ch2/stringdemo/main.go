package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// var s1 string = "mysql ... -u root -p 123"
// var s2 string
// var s3 string = "北京 上海 广州 深圳"

func main() {
	// uindex := strings.Index(s1, "-u")
	// pindex := strings.Index(s1, "-p")
	// username := s1[uindex+3 : pindex-1]
	// password := s1[pindex+3:]
	// if username == "root" && password == "123" {
	// 	fmt.Println("正确")
	// } else {
	// 	fmt.Println("错误")
	// }
	// fmt.Println("请输入：")
	// fmt.Scan(&s2)
	// if strings.HasPrefix(s2, "A") || strings.HasPrefix(s2, "a") {
	// 	fmt.Println("true")

	// } else {
	// 	fmt.Println("false")
	// }

	// s3 := strings.Split(s3, " ")
	// fmt.Println(strings.Join(s3, ","))

	// 	类型转换
	// 整型转换
	var x int8 = 10
	var y int64 = 20
	fmt.Println(x + int8(y))

	// 字符串与整型之间转换： strconv
	var ageStr = "32"
	var age, _ = strconv.Atoi(ageStr)
	// fmt.Println("err", err) //<nil>
	fmt.Println(age, reflect.TypeOf(age))

	price := 100
	priceStr := strconv.Itoa(price)
	fmt.Println(priceStr, reflect.TypeOf(priceStr))

	// strconv Parse系列函数
	// 将数字字符串转换为
	ret, _ := strconv.ParseInt("28", 10, 8)
	fmt.Println(ret, reflect.TypeOf(ret))
}
