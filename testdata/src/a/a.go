package a

type MyInt int
type MyMyInt MyInt
type I interface {
	string | ~int | bool
}

// error: empty
// type A interface {
// }

// error: contains method
// type B interface {
// 	string | int | float64
// 	Equal() bool
// }

// type D interface {
// 	string | ~int
// 	MyInt
// }

type E interface {
	int | string
	float64
}

// Bug?
// If a single type with ~ is included, it is interpreted as *types.Union.
// With the current go/types package, the following two are impossible to distinguish.
// So, it is necessary to traverse ASTs using *ast.BinaryExpr, etc. in the go/ast package.
// type F interface {
// 	int | string
// 	~float64
// }
// type F interface {
// 	int | string | ~float64
// }

// type G interface {
// 	string | float64
// 	~int
// }

// type MyInt int
// type H interface {
// 	~int | ~string
// 	MyInt
// }
