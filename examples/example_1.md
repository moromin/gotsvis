- target source file
```go
type I interface {
	string | ~int | bool
}
```

- output
```bash
TypeSet: "I"
---------------------
string: 10
~int: 100
bool: 10
string ∩ ~int: 0
string ∩ bool: 0
~int ∩ bool: 0
string ∩ ~int ∩ bool: 0
```

- venn diagram ([venn.js](https://github.com/benfred/venn.js/))
![Example 1](https://github.com/moromin/gotsvis/blob/dataset/images/vennjs_OR_exapmle.png)
