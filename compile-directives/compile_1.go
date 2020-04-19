package main

import (
	"fmt"
)

/*

//go:noinline
  noinline 顾名思义，不要内联。
  Inline 内联
	Inline，是在编译期间发生的，将函数调用调用处替换为被调用函数主体的一种编译器优化手段。Wiki：Inline 定义 https://en.wikipedia.org/wiki/Inline_expansion
	使用 Inline 有一些优势，同样也有一些问题。
	优势：
	减少函数调用的开销，提高执行速度。
	复制后的更大函数体为其他编译优化带来可能性，如 过程间优化 https://en.wikipedia.org/wiki/Interprocedural_optimization
	消除分支，并改善空间局部性和指令顺序性，同样可以提高性能。
	问题：
	代码复制带来的空间增长。
	如果有大量重复代码，反而会降低缓存命中率，尤其对 CPU 缓存是致命的。
	所以，在实际使用中，对于是否使用内联，要谨慎考虑，并做好平衡，以使它发挥最大的作用。
	简单来说，对于短小而且工作较少的函数，使用内联是有效益的。
*/
/*
执行 GOOS=linux GOARCH=386 go tool compile -S compile_1.go > compile_1.S
这个地方会出错，请到 ex 文件夹下面测试
*/
func appendStr(word string) string {
	return "new " + word
}

//go:noinline
func appendStr2(word string) string {
	return "new " + word
}

func main() {
	fmt.Println(appendStr("aa"))

	fmt.Println(appendStr2("bb"))
}
