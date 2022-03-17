package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/types"
	"os"

	"github.com/moromin/gotsvis"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"
)

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("パターンを指定してください")
	}

	cfg := &packages.Config{
		Mode: packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}
	pkgs, err := packages.Load(cfg, os.Args[1:]...)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		if err := analyze(pkg); err != nil {
			return err
		}
	}

	return nil
}

func analyze(pkg *packages.Package) error {
	inspect := inspector.New(pkg.Syntax)

	nodeFilter := []ast.Node{
		(*ast.InterfaceType)(nil),
		(*ast.TypeSpec)(nil),
	}

	var title string

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			if n.Type != nil {
				title = n.Name.Name
			}
		case *ast.InterfaceType:
			m := pkg.TypesInfo.TypeOf(n).(*types.Interface)

			// check error case
			if m.Empty() {
				fmt.Fprintf(os.Stderr, "interface %q is empty...\n", title)
			} else if m.NumMethods() != 0 {
				fmt.Fprintf(os.Stderr, "interface %q is not type set...\n", title)
			}
			if n.Methods == nil {
				return
			}

			res := make([]types.Type, 0)
			for _, e := range n.Methods.List {
				typ := pkg.TypesInfo.TypeOf(e.Type)
				res = append(res, typ)
			}
			gotsvis.Venn(title, res)
		}
	})

	return nil
}
