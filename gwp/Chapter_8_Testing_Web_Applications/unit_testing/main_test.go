package main // 测试文件与被测试的源代码文件位于同一个包内

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json") // 调用被测试的函数
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 { // 检查结果是否和预期一样，如果不一样就显示一条出错信息
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong content, was expecting 'Hello World!' but got", post.Content)
	}
}
func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now") // 暂时跳过对编码函数的测试。
}
func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}
