package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)
	writer = httptest.NewRecorder()
}
func TestHandleGet(t *testing.T) {
	mux := http.NewServeMux()               // 创建一个用于测试的多路复用器
	mux.HandleFunc("/post/", handleRequest) // 绑定想要测试的处理其

	writer := httptest.NewRecorder()                     // 创建记录器，用于获取服务器返回的HTTP响应
	request, _ := http.NewRequest("GET", "/post/1", nil) //为被测试的处理器创建相应的请求
	mux.ServeHTTP(writer, request)                       // 向被测试的处理器发送请求

	if writer.Code != 200 { // 对记录器记载响应结果进行检查
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}
}
func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
