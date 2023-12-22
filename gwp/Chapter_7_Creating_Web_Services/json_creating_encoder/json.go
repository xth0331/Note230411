package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Post struct { // 创建结构体并向里面填充数据
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:      1,
		Content: "Hello World!",
		Author: Author{
			Id:   2,
			Name: "Sau sheong",
		},
		Comments: []Comment{
			Comment{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			Comment{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}
	// 创建用于存储数据的json文件
	jsonFile, err := os.Create("gwp/Chapter_7_Creating_Web_Services/json_creating_encoder/post.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	encoder := json.NewEncoder(jsonFile) // 根据给定的JSON文件创建出相应的编码器
	err = encoder.Encode(&post)
	if err != nil { // 把结构编码写到JSON文件中。
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}
