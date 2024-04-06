package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// 获取指定的帖子
// GET /post/1
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id) // 从数据库里获取数据，并将其填充到Post结构中
	if err != nil {
		return

	}
	output, err := json.MarshalIndent(&post, "", "\t\t") // 把Post结构封装为JSON字符串
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json") // 把JSON数据写入ResponseWriter
	w.Write(output)
	return
}

// 创建新的帖子
// POST /post/
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len) // 读取请求主体，并将其存储在字节切片中；创建一个字节切片。
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post) // 切片存储的数据解封至Post结构体。
	err = post.create()         // 创建数据库记录。
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// 更新指定的帖子
// PUT /post/1
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id) // 从数据库里获取指定帖子的数据，并将其填充至Post结构
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)           // 从请求主体中读取JSON数据
	json.Unmarshal(body, &post) // JSON数据解封至Post结构
	err = post.update()         // 对数据库进行更新
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// 删除指定的帖子
// DELETE /post/1
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id) // 从数据库获取指定帖子的数据，并将其填充至Post结构
	if err != nil {
		return
	}
	err = post.delete() // 从数据库删除这个帖子
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return

}

// 多路复用器负责将请求转发给正确的处理器函数

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func main() {
	server := http.Server{
		Addr: ":8080"}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}
