- target source file
```go
type I interface {
	int | string
	~float64
}
```

- output
```bash
TypeSet: "I"
Empty set.
```

- venn diagram

	nil
