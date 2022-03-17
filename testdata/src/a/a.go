package a

// error: empty
// type A interface {
// }

// error: contains method
// type B interface {
// 	string | int | float64
// 	Equal() bool
// }

type C interface {
	~int | string
	string
}

type D interface {
	int | string
	float64
}

type E interface {
	int | string
	~float64
}

type F interface {
	string | int
	~int
}

type G interface {
	~int | ~string
	string
}

type H interface {
	~int | ~string
	~string
}

// TODO: support the following
// type MyInt int
// type Todo interface {
// 	string | ~int | bool
// 	MyInt
// }
