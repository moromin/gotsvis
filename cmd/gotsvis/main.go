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

	// AST dump
	// for i, _ := range pkgs {
	// 	fset := token.NewFileSet()
	// 	f, err := parser.ParseFile(fset, os.Args[i+1], nil, 0)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}
	// 	ast.Print(fset, f)
	// }

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
				// fmt.Printf("Type set: %s\n", n.Name.Name)
			}
		case *ast.InterfaceType:
			if n.Methods == nil {
				return
			}
			res := make([]types.Type, 0)
			for _, e := range n.Methods.List {
				typ := pkg.TypesInfo.TypeOf(e.Type)
				res = append(res, typ)
				// fmt.Printf("\t%[1]T %[1]v\n", typ.String())
				// fmt.Println()
				// set := make([]types.Type, 0)
				// switch typ := typ.(type) {
				// case *types.Union:
				// fmt.Printf("\t%[1]T %[1]v\n", typ.String())
				// for i := 0; i < typ.Len(); i++ {
				// set = append(set, typ.Term(i))
				// term := typ.Term(i)
				// fmt.Printf("\t\t%[1]T %[1]v\n", typ.Term(i))
				// if term.Tilde() {
				// } else {
				// 	fmt.Printf("\t\t%[1]T %[1]v\n", term.Type())
				// }
				// }
				// res = append(res, set)
				// default:
				// res = append(res, []string{typ.String()})
				// fmt.Printf("\t%[1]T %[1]v\n", typ)
				// }
			}
			gotsvis.Venn(title, res)
		}
	})

	return nil
}
