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

	ast.Print(pass.Fset, pass.Files[0])

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			fmt.Println("TypeSpec")
			typ := pass.TypesInfo.TypeOf(n.Type)
			fmt.Println(typ)
		}
	})

	return nil, nil
}
