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

	res := make([][]string, 0)
	for _, field := range list {
		set := make([]string, 0)
		switch n := field.Type.(type) {
		case *ast.BinaryExpr:
			set = append(set, binaryExprToSlice(n)...)
		case *ast.UnaryExpr:
			ident, _ := n.X.(*ast.Ident)
			set = append(set, fmt.Sprintf("%s%s", n.Op, ident.Name))
		case *ast.Ident:
			set = append(set, n.Name)
		// TODO: handle FuncType
		case *ast.FuncType:
			set = append(set, "some method")
		}
		res = append(res, set)
	}
	pass.Reportf(node.Pos(), "%v", res)
}
