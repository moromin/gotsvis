# gotsvis
`gotsvis` visualize TypeSet.

The concept of typesets was added in Go 1.18.
Type sets is a new concept for describing the conditions under which a type "implements" an interface.

`gotsvis` provides dataset that can be used to draw a Venn diagram.

## Usage
```bash
git clone https://github.com/moromin/gotsvis
go run ./cmd/gotsvis/main.go <target_source_file>
```

## Example
- target source file
```go
type I interface {
	string | ~int | bool
}
```

- output
```bash
~int: 100
bool: 10
string: 10
~int ∩ bool: 0
string ∩ ~int: 0
string ∩ bool: 0
string ∩ ~int ∩ bool: 0
```

![Example 1](https://github.com/moromin/gotsvis/blob/dataset/images/vennjs_OR_exapmle.png)
