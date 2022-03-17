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

type D interface {
	~int | string
	string
}

type E interface {
	int | string
	float64
}

// n.Len() is different
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
