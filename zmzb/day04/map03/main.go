package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var data = make(map[string][]string)
	// data["beijing"] = []string{"朝阳", "海淀", "昌平"}
	// data["shandong"] = []string{"济南", "青岛", "威海"}
	// data["hebei"] = []string{"唐山", "石家庄", "保定"}
	// // fmt.Println(data)
	// // // 查询河北第二个城市
	// // fmt.Println(data["hebei"][1])
	// //遍历每一个省份以及对应的城市名称
	// // for proStr, citysSlice := range data {
	// // 	fmt.Println(proStr)
	// // 	for i, city := range citysSlice {
	// // 		fmt.Printf("%d:%s", i, city)
	// // 	}
	// // 	fmt.Println()
	// // }
	// // 循环打印每个省份的名字和城市数量
	// // for proStr, citySlice := range data {
	// // 	fmt.Printf("省份:%s,城市数:%d\n", proStr, len(citySlice))
	// // }
	// // 添加一个新的省份和城市
	// // data["heilongjiang"] = []string{"哈尔滨", "齐齐哈尔", "佳木斯"}
	// // 删除北京
	// delete(data, "beijing")
	// fmt.Println(data)
	// 案例2
	// info := map[int]map[string]string{1001: {"name": "yuan", "age": "23"}, 1002: {"name": "alvin", "age": "33"}}
	// fmt.Println(info[1002]["age"])
	// for num, stuinfo := range info {
	// 	fmt.Printf("学号:%d,姓名:%s,年龄:%s\n", num, stuinfo["name"], stuinfo["age"])
	// }
	// info[1003] = map[string]string{"name": "xie", "age": "30"}
	// delete(info, 1001)
	// fmt.Println(info)

	// 打印学号为1002的学生的年龄
	// 循环打印每个学员的学号，姓名，年龄
	// 添加一个新的学员
	// 删除1001的学生
	// 案例3
	stus := []map[string]string{{"name": "yuan", "age": "23"}, {"name": "rain", "age": "22"}, {"age": "32", "name": "eric"}}
	// fmt.Println(stus[1]["name"])
	for _, stu := range stus {
		fmt.Printf("姓名:%s,年龄:%s\n", stu["name"], stu["age"])
		if stu["name"] == "rain" {
			ageInt, _ := strconv.Atoi(stu["age"])
			ageInt += 1
			stu["age"] = strconv.Itoa(ageInt)
		}
	}
	// stus = append(stus, map[string]string{"name": "xie", "age": "30"})
	fmt.Println(stus)

	// 打印第二个学生的姓名
	// 循环打印每一个学生的姓名和年龄
	// 添加一个新的学生map
	// 删除一个学生map
	// 将姓名为rain的学生的年龄自加一岁

}
