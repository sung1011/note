package main

import (
	"testing"
	"time"
)

/*
go test -benchmem -run=^$ -bench ^BenchmarkFib$ scr/basic -v

// 2核 和 4核
go test -benchmem -run=^$ -bench ^BenchmarkFib$ scr/basic -v -cpu=2,4

// 执行5秒
go test -benchmem -run=^$ -bench ^BenchmarkFib$ scr/basic -v -benchtime=5s

// 执行5000次
go test -benchmem -run=^$ -bench ^BenchmarkFib$ scr/basic -v -benchtime=5000x

// 重复3次
go test -benchmem -run=^$ -bench ^BenchmarkFib$ scr/basic -v -count=3
*/

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // 关注每次执行的纳秒情况 xx ns/op
	}
}

func BenchmarkSliceNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceNoCap() // 关注内存分配次数 20 allocs/op
	}
}
func BenchmarkSliceWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceWithCap() // 关注内存分配次数 0 allocs/op
	}
}

func BenchmarkFibResetTimer(b *testing.B) {
	time.Sleep(5 * time.Second) // 模拟耗时的准备工作
	b.ResetTimer()              // 重置计时器，忽略前面的准备时间; (实测未生效)
	for n := 0; n < b.N; n++ {
		fib(10)
	}
}

func BenchmarkFibStopTimer(b *testing.B) {
	b.StopTimer()               // 暂停计时 (实测未生效)
	time.Sleep(5 * time.Second) // 每次函数执行前的准备工作
	b.StartTimer()              // 继续计时
	for n := 0; n < b.N; n++ {
		fib(10)
	}
}

// --------------------------------------------

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func sliceNoCap() {
	sl := make([]int, 0, 0)
	for i := 0; i < 10000; i++ {
		sl = append(sl, i)
	}
}

func sliceWithCap() {
	sl := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		sl = append(sl, i)
	}
}
