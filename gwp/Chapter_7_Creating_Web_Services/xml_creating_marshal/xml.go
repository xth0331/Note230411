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
		Id:      "1",
		Content: "Hello World!", // 创建结构并向里面填充数据
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}
	//output, err := xml.Marshal(&post)
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil { // 把结构封装为由字节切片组成的XML数据。
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	err = os.WriteFile("gwp/Chapter_7_Creating_Web_Services/xml_creating_marshal/post.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}
