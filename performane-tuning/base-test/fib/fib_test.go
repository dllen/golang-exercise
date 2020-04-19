package fib

import (
	"testing"
)

// 此函数计算斐波那契数列中第 N 个数字
func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

func Fib2(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

//go test -bench=.
//go test -bench=. -cpu=1,2,4
//go test -bench=. -benchtime=10s
//go test -bench=. -count=10
func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20) // 运行 Fib 函数 N 次
	}
}
