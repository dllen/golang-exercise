package main

/*
源代码扫描完成后，扫描结果将被传递给语法分析器。语法分析是编译的一个阶段，它将 token 转换为 抽象语法树（AST）。
AST 是源代码的结构化表示。在 AST 中，我们将能够看到程序结构，比如函数和常量声明。
我们使用 go/parser 和 go/ast 来打印完整的 AST
*/
/*
分析这个输出，在 Decls 字段中，包含了代码中所有的声明，例如导入、常量、变量和函数。在本例中，我们只有两个：导入fmt包 和 主函数。
*/
import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	src := []byte(`
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
`)

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	ast.Print(fset, file)
}
