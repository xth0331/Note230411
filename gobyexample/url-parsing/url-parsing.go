package main

import (
	"fmt"
	"net"
	"net/url"
)

// URL提供统一资源定位方式。这里展示了在Go中时如何解析URL的。
func main() {
	// 我们将解析这个URL示例，它包含了一个scheme、认证信息、主机名、认证信息、主机名、端口、路径、查询参数以及查询片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	// 解析这个URL并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)

	}
	// 直接访问scheme
	fmt.Println(u.Scheme)
	// User包含了所有的认证信息，这里调用Username和Password来获取单独的值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	// Host包含主机名和端口号（如果存在）。使用SplitHostPort提取它们。
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	// 这里我们提取路径和#后面的查询片段信息。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	//要得到字符串中的 k=v 这种格式的查询参数，可以使用 RawQuery 函数。 你也可以将查询参数解析为一个 map。
	//已解析的查询参数 map 以查询字符串为键， 已解析的查询参数会从字符串映射到到字符串的切片， 因此如果您只想要第一个值，则索引为 [0]。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
