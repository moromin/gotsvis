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

	// Map defined type -> underlying type
	du := make(map[string]string)

	inspect.Preorder(nil, func(n ast.Node) {
		s := pkg.TypesInfo.Scopes[n]
		if s == nil {
			return
		}
		for _, name := range s.Parent().Names() {
			if typ, ok := s.Parent().Lookup(name).Type().(*types.Named); ok {
				du[name] = typ.Underlying().String()
			}
		}
	})

	// Search type set
	var title string

	nodeFilter := []ast.Node{
		(*ast.InterfaceType)(nil),
		(*ast.TypeSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			if n.Type != nil {
				title = n.Name.Name
			}
		case *ast.InterfaceType:
			ifn := pkg.TypesInfo.TypeOf(n).(*types.Interface)

			// check error case
			if ifn.Empty() {
				fmt.Fprintf(os.Stderr, "interface %q is empty...\n", title)
			} else if ifn.NumMethods() != 0 {
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
			gotsvis.Venn(title, res, du)
		}
	})

	return nil
}
