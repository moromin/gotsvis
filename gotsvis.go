package gotsvis

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "gotsvis visualize type set"

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
	// fmt.Println(pass.TypesInfo)

	inspect.Preorder(nil, func(n ast.Node) {
		if n, ok := n.(*ast.TypeSpec); ok {
			if typ, ok := n.Type.(*ast.InterfaceType); ok {
				printParams(pass, typ, n)
			}
		}
	})

	return nil, nil
}

func printParams(pass *analysis.Pass, typ *ast.InterfaceType, node *ast.TypeSpec) {
	methods := (*typ).Methods
	list := (*methods).List

	res := make([]string, 0)
	for _, field := range list {
		switch n := field.Type.(type) {
		case *ast.BinaryExpr:
			res = append(res, binaryExprToSlice(n)...)
		case *ast.UnaryExpr:
			ident, _ := n.X.(*ast.Ident)
			res = append(res, fmt.Sprintf("%s%s", n.Op, ident.Name))
		case *ast.Ident:
			res = append(res, n.Name)
		// TODO: handle FuncType
		case *ast.FuncType:
			res = append(res, "some method")
		}
	}
	pass.Reportf(node.Pos(), "%v", res)
}
