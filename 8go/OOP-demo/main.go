package main

import (
	"fmt"
	"reflect"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}
func main() {
	s := NewServer("127.0.0.1", 8080)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
