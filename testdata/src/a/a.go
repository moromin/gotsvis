package a

type A interface {
	string
}

type B interface {
	string | int | float64
}

type C interface {
	~string | int
}

type D interface {
	string | int | float64 | bool
}

type E interface {
	int | string
	float64
}
