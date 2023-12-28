package main

import "testing"

// benchmarking the decode function
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ { // 循环执行解码函数，以便对其进行b.N次基准测试。
		decode("post.json")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}
