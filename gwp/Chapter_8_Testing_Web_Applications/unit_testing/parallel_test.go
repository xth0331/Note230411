package main

import (
	"testing"
	"time"
)

func TestParallel_1(t *testing.T) { // 模拟需要耗时一秒运行的任务
	t.Parallel() // 调用Parallel函数，以并行的方式运行测试用例
	time.Sleep(1 * time.Second)
}
func TestParallel_2(t *testing.T) { // 模拟需要耗时2秒的任务
	t.Parallel()
	time.Sleep(2 * time.Second)
}
func TestParallel_3(t *testing.T) { // 模拟需要耗时3秒的任务
	t.Parallel()
	time.Sleep(3 * time.Second)
}
