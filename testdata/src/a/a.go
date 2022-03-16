package a

func f1[T any](v T) {
	print(v)
}

type A[T ~bool] []T
type B[T ~int] []T

type C interface {
	~string | int
}
