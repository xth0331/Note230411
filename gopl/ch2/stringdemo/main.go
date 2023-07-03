package main

import "fmt"

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
	// var x int8 = 10
	// var y int64 = 20
	// fmt.Println(x + int8(y))

	// // 字符串与整型之间转换： strconv
	// var ageStr = "32"
	// var age, _ = strconv.Atoi(ageStr)
	// // fmt.Println("err", err) //<nil>
	// fmt.Println(age, reflect.TypeOf(age))

	// price := 100
	// priceStr := strconv.Itoa(price)
	// fmt.Println(priceStr, reflect.TypeOf(priceStr))

	// // strconv Parse系列函数
	// // 将数字字符串转换为
	// ret, _ := strconv.ParseInt("28", 10, 8)
	// fmt.Println(ret, reflect.TypeOf(ret))
	// // 将字符串转换浮点型
	// ret2, _ := strconv.ParseFloat("3.1415", 64)
	// fmt.Println(ret2, reflect.TypeOf(ret2))

	// // 字符串转换成布尔
	// b0, _ := strconv.ParseBool("t")
	// fmt.Println(b0)
	// b1, _ := strconv.ParseBool("f")
	// fmt.Println(b1)
	// b2, _ := strconv.ParseBool("0")
	// fmt.Println(b2)
	// b3, _ := strconv.ParseBool("1")
	// fmt.Println(b3)
	// 输入函数IO
	//var name string
	//var age int
	//fmt.Scan(&name,&age) // 等待用户输入
	//fmt.Println(name)
	//fmt.Println("end")
	//var birthday string
	//fmt.Println("请输入您的生日，使用-分隔，例如1990-3-1")
	//fmt.Scanln(&birthday)
	//birthdaySplit := strings.Split(birthday, "-")
	//birMonth, _ := strconv.Atoi(birthdaySplit[1])
	//birDay, _ := strconv.Atoi(birthdaySplit[2])
	//switch birMonth {
	//case 1:
	//	if birDay <= 19 {
	//		fmt.Println("摩羯座")
	//	} else {
	//		fmt.Println("水瓶座")
	//	}
	//case 2:
	//	if birDay <= 18 {
	//		fmt.Println("水瓶座")
	//	} else {
	//		fmt.Println("双鱼座")
	//	}
	//case 3:
	//	if birDay <= 20 {
	//		fmt.Println("双鱼座")
	//	} else {
	//		fmt.Println("白羊座")
	//	}
	//case 4:
	//	if birDay <= 19 {
	//		fmt.Println("白羊座")
	//	} else {
	//		fmt.Println("金牛座")
	//	}
	//case 5:
	//	if birDay <= 20 {
	//		fmt.Println("金牛座")
	//	} else {
	//		fmt.Println("双子座")
	//	}
	//case 6:
	//	if birDay <= 21 {
	//		fmt.Println("双子座")
	//	} else {
	//		fmt.Println("巨蟹座")
	//	}
	//case 7:
	//	if birDay <= 22 {
	//		fmt.Println("巨蟹座")
	//	} else {
	//		fmt.Println("狮子座")
	//	}
	//case 8:
	//	if birDay <= 22 {
	//		fmt.Println("狮子座")
	//	} else {
	//		fmt.Println("处女座")
	//	}
	//case 9:
	//	if birDay <= 22 {
	//		fmt.Println("处女座")
	//	} else {
	//		fmt.Println("天秤座")
	//	}
	//case 10:
	//	if birDay <= 23 {
	//		fmt.Println("天秤座")
	//	} else {
	//		fmt.Println("天蝎座")
	//	}
	//case 11:
	//	if birDay <= 22 {
	//		fmt.Println("天蝎座")
	//	} else {
	//		fmt.Println("射手座")
	//	}
	//case 12:
	//	if birDay <= 21 {
	//		fmt.Println("射手座")
	//	} else {
	//		fmt.Println("摩羯座")
	//	}
	//}
	var sum int
	for count := 1; count <= 100; count++ {
		sum += count
	}
	fmt.Println(sum)

}
