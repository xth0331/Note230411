package main

import (
	"encoding/xml"
	"fmt"
)

// Plant 结构将被映射到 XML 。 与 JSON 示例类似，字段标签包含用于编码器和解码器的指令。
// 这里我们使用了 XML 包的一些特性： XMLName 字段名称规定了该结构的 XML 元素名称； id,attrr 表示 Id 字段是一个 XML 属性，而不是嵌套元素。

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v,name=%v,origin=%v", p.Id, p.Name, p.Origin)
}
func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	//传入我们声明了 XML 的 Plant 类型。 使用 MarshalIndent 生成可读性更好的输出结果。
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	// 明确的为输出结果添加一个通用的XML头部信息。
	fmt.Println(xml.Header + string(out))
	// 使用Unmarshal将XML格式字节流到Plant结构。如果XML格式错误或无法映射到Plant
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}
	// parent>child>plant字段标签告诉编码器，将Plants中的元素嵌套在<parent><child>里面。
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"plants>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))

}
