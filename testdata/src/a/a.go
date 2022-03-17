package a

// type A interface {
// 	string
// }
type MyInt int
type I interface {
	string | ~int | bool
}

// type B interface {
// 	string | int | float64
// }

// type D interface {
// 	string | int | float64 | bool
// }

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
