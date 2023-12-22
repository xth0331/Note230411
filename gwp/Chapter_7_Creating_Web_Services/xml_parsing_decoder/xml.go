package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName  string    `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("gwp/Chapter_7_Creating_Web_Services/xml_parsing_decoder/post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	// 根据给定的XML数据生成相应的解码器
	decoder := xml.NewDecoder(xmlFile)
	for { // 每迭代一次解码器中的所有XML数据
		t, err := decoder.Token() // 每进行一次迭代，就从解码器里面获取一个token
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}
		switch se := t.(type) { // 检查token的类型
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se) // 将XML数据解码至结构
				fmt.Println(comment)
			}
		}
	}
}
