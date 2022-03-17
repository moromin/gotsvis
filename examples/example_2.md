- target source file
```go
type I interface {
	~int | ~string
	string
}
```

- output
```bash
TypeSet: "I"
---------------------
~int: 100
~string: 100
string: 10
~int ∩ ~string: 0
~int ∩ string: 0
~string ∩ string: 10
~int ∩ ~string ∩ string: 0
```

- venn diagram ([venn.js](https://github.com/benfred/venn.js/))
![Example 2](https://github.com/moromin/gotsvis/blob/main/images/vennjs_OR_AND_example.png)
