package gotsvis

import (
	"fmt"
	"go/ast"
)

type visitorFunc func(node ast.Node) (w ast.Visitor)

func (f visitorFunc) Visit(node ast.Node) (w ast.Visitor) {
	return f(node)
}

func binaryExprToSlice(root *ast.BinaryExpr) []string {
	res := make([]string, 0)

	var v2 visitorFunc
	v2 = visitorFunc(func(node ast.Node) (w ast.Visitor) {
		if node == nil {
			return nil
		}
		switch n := node.(type) {
		case *ast.Ident:
			res = append(res, n.Name)
		case *ast.UnaryExpr:
			ident, _ := n.X.(*ast.Ident)
			res = append(res, fmt.Sprintf("%s%s", n.Op, ident.Name))
		case *ast.BinaryExpr:
			ast.Walk(v2, n.X)
			ast.Walk(v2, n.Y)
		}
		return nil
	})
	ast.Walk(v2, root)

	return res
}
