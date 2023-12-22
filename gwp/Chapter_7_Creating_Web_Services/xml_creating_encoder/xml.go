package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1", // 创建结构并向里面填充数据
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}
	xmlFile, err := os.Create("gwp/Chapter_7_Creating_Web_Services/xml_creating_encoder/post.xml") // 创建用于存储数据的XML文件
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(xmlFile) // 根据给定的XML文件，创建出相应的编码器。
	encoder.Indent("", "\t")
	err = encoder.Encode(&post) // 把结构编码至文件
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}
