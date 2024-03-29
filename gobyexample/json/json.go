package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 下面我们将使用者两个结构体来演示自定义类型的编码和解码。
type response1 struct {
	Page   int
	Fruits []string
}

// 只有可导出的字段才会被JSON编码/解码。必须以大写字母开头的字段才是可导出的。
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// 首先我们来看一下基本数据类型到JSON字符串的编码过程。这是一些原子值的例子。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 这是一些切片和map编码JSON数组和对象的例子。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	// JSON 包可以自动的编码你的自定义类型。 编码的输出只包含可导出的字段，并且默认使用字段名作为 JSON 数据的键名。
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	// 你可以给结构体字段声明标签来自定义编码的JSON数据的键名。上面Response2定义，就是这种标签的一个例子。
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	// 现在来看看将JSON数据解码为Go值的过程。这是一个普通数据结构解码例子。
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//我们需要提供一个JSON包可以存放解码数据的变量。这里的map[string]interface{}是一个键为string，值为任意值的map。
	var dat map[string]interface{}
	// 这是实际的解码，以及相关错误的检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 为了使用解码map中的值，我们需要将它们进行适当的类型转换。例如，这里我们将num的值转换成float64类型。
	num := dat["num"].(float64)
	fmt.Println(num)
	// 访问嵌套的值需要一系列的转化。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	// 我们还可以将 JSON 解码为自定义数据类型。 这样做的好处是，可以为我们的程序增加附加的类型安全性， 并在访问解码后的数据时不需要类型断言。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
	// 在上面例子的标准输出上， 我们总是使用 byte和 string 作为数据和 JSON 表示形式之间的中介。
	// 当然，我们也可以像 os.Stdout 一样直接将 JSON 编码流传输到 os.Writer 甚至 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
