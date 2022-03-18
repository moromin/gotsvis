# gotsvis
`gotsvis` visualize TypeSet.

The concept of typesets was added in Go 1.18.
Type sets is a new concept for describing the conditions under which a type "implements" an interface.

`gotsvis` doesn't support interface that has duplicate type set or method set.

`gotsvis` provides dataset that can be used to draw a Venn diagram.

## Usage
```bash
git clone https://github.com/moromin/gotsvis
go run ./cmd/gotsvis/main.go <target_source_file>
```


## Drawing Venn diagram tools
- Javascript : [venn.js](https://github.com/benfred/venn.js)
- Google Charts : [Venn Charts](https://developers.google.com/chart/image/docs/gallery/venn_charts)(deprecated)
- Python : [matplotlib-venn](https://pypi.org/project/matplotlib-venn/)
- R : [gplots/venn](https://www.rdocumentation.org/packages/gplots/versions/3.1.1/topics/venn)


## Example
- [examples](https://github.com/moromin/gotsvis/tree/main/examples)
