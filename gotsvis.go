package gotsvis

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "gotsvis is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "gotsvis",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// ast.Print(pass.Fset, pass.Files[0])

	inspect.Preorder(nil, func(n ast.Node) {
		if n, ok := n.(*ast.TypeSpec); ok {
			printParams(n)
		}
	})

	return nil, nil
}

func printParams(node *ast.TypeSpec) {
	typ, ok := node.Type.(*ast.InterfaceType)
	if !ok {
		return
	}

	fmt.Println(node.Name)

	methods := (*typ).Methods
	list := (*methods).List

	// confirm node
	res := make([]string, 0)
	for _, field := range list {
		switch n := field.Type.(type) {
		case *ast.BinaryExpr:
			res = append(res, binaryExprToSlice(n)...)
		case *ast.UnaryExpr:
			ident, _ := n.X.(*ast.Ident)
			// fmt.Printf("%s%s\n", n.Op, ident.Name)
			res = append(res, fmt.Sprintf("%s%s", n.Op, ident.Name))
		case *ast.Ident:
			// fmt.Println(n.Name)
			res = append(res, n.Name)
		}
	}
	fmt.Println(res)
}
