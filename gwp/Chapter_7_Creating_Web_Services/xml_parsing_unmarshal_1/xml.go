package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("gwp/Chapter_7_Creating_Web_Services/xml_parsing_unmarshal_1/post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return

	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
