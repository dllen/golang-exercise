package main

/*
AST 包含了许多信息，我们不仅可以分析出以上结论，还可以进一步检查 AST 并查找文件中的所有函数调用。下面，我们将使用 go/ast 包中的 Inspect 函数来递归地遍历树，并分析所有节点的信息。
*/
import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
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
		fmt.Println(err)
	}
	//查找所有节点以及它们是否为 *ast.CallExpr 类型，上面也说过这种类型是函数调用。如果是，则使用 go/printer 包打印 Fun 中存在的函数的名称。
	ast.Inspect(file, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		printer.Fprint(os.Stdout, fset, call.Fun)

		return false
	})
}
