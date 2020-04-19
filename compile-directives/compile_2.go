package main

/*
norace 的作用是：跳过竞态检测
我们知道，在多线程程序中，难免会出现数据竞争，正常情况下，当编译器检测到有数据竞争，就会给出提示。

执行 go run -race compile_2.go 利用 -race 来使编译器报告数据竞争问题
*/

var sum int

//go:norace
func add() int {
	sum++
	return sum
}

func add2() int {
	sum++
	return sum
}

func main() {
	//检查没有警告
	go add()
	go add()

	//检查会有警告
	go add2()
	go add2()
}
