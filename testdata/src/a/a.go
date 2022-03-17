package a

type A interface { // want "[string]"
	string
}
type C interface { // want "[~string int]"
	~string | int
}

// type B interface { // want "[string int float64]"
// 	string | int | float64
// }

// type D interface { // want "[string int float64 bool]"
// 	string | int | float64 | bool
// }

// type E interface { // want "[int string float4]"
// 	int | string
// 	float64
// }

// type F interface { // want "[int string float64]"
// 	int | string
// 	float64
// }

// type G interface {
// 	string | float64
// 	~int
// 	Equal() bool
// }

// type MyInt int
// type H interface {
// 	~int | string
// 	int | ~string
// }
