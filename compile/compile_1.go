package main

/*
扫描代码，进行词法分析
任何编译器的第一步都是将源代码文本分解成 token，由扫描程序（也称为词法分析器）完成。token 可以是关键字，字符串，变量名，函数名等等。每一个有效的词都由 token 表示。
在 Go 中，我们写在代码上的 "package"，"main"，"func" 这些都是 token。

token 由代码中的位置，类型和原始文本组成。我们可以使用 go/scanner 和 go/token 包在 Go 程序中自己执行扫描程序。这意味着我们可以像编译器那样扫描检视自己的代码。
下面，我们将通过一个打印 Hello World 的示例来展示 token。
*/
/*
以第一行为例分析这个输出，第一列 2:1 表示扫描到了源代码第二行第一个字符，第二列 package 表示 token 是 package，第三列 "package" 表示源代码文本。
我们可以看到在 Scanner 执行过程中将 \n 换行符标记成了 ; 分号，像在 C 语言中是用分号表示一行结束的。这就解释了为什么 Go 不需要分号：它们是在词法分析阶段由 Scanner 智能地解释的。
*/
import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	src := []byte(`
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
`)

	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, 0)

	for {
		pos, tok, lit := s.Scan()
		fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

		if tok == token.EOF {
			break
		}
	}
}
