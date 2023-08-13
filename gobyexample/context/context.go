package main

import (
	"fmt"
	"net/http"
	"time"
)

// HTTP 服务器对于演示 context.Context 的用法很有用的， context.Context 被用于控制 cancel。 Context 跨 API 边界和协程携带了：deadline、取消信号以及其他请求范围的值。

func hello(w http.ResponseWriter, req *http.Request) {
	// net/http机制为每个请求创建来一个context.Context，并且可以通过Context()方法获取它。
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")
	// 扽带几秒钟，然后再将回复发送到客户端。只可以模拟服务器正在执行的某些工作。
	// 在工作时，请密切关注context的Done()通道的信号，一旦收到该信号，表明我们应该取消工作并尽快返回。
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// context的Err()方法返回来一个错误，该错误说明了Done通道关闭的原因。
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
func main() {
	// 跟前面一样，我们在/hello路由上注册handler，然后开始提供服务。
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
