package a

type MyInt int
type I interface {
	string | ~int | bool
}

type A interface {
}

type B interface {
	string | int | float64
	Equal() bool
}

type D interface {
	string | ~int
	MyInt
}

// type E interface {
// 	int | string
// 	float64
// }

// type F interface {
// 	int | string
// 	float64
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
