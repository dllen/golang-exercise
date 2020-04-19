package main

/*
为了显示生成的 SSA，我们需要将 GOSSAFUNC 环境变量设置为我们想要跟踪的函数，在本例中为main 函数。我们还需要将 -S 标识传递给编译器，这样它就会打印代码并创建一个HTML文件。我们还将编译Linux 64位的文件，以确保机器代码与您在这里看到的相同。
在终端执行下面的命令：

GOSSAFUNC=main GOOS=linux GOARCH=amd64 go build -gcflags -S compile_4.go

会在终端打印出所有的 SSA，同时也会生成一个交互式的 ssa.html 文件，我们用浏览器打开它。
*/

import "fmt"

func main() {
	fmt.Println(2)
}
