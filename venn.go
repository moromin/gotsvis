package gotsvis

import (
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/exp/slices"
)

func Venn(title string, s []types.Type, du map[string]string) {
	typeSlice := []string{}
	subsets := [][]string{}
	numOfSubset := make(map[string]int)

	for _, elem := range s {
		switch n := elem.(type) {
		case *types.Union:
			// TODO: ~*** intersection other, (ex. int | string; ~int, *types.Union.Len() == 1)
			if n.Len() == 1 {
				typName := n.String()
				if !slices.Contains(typeSlice, typName) {
					typeSlice = append(typeSlice, typName)
					subsets = getSubsetCombination(typeSlice)
				}
				numOfSubset = intersectionSubset(subsets, numOfSubset, typName)
			} else {
				for i := 0; i < n.Len(); i++ {
					typeSlice = append(typeSlice, n.Term(i).String())
				}
				subsets = getSubsetCombination(typeSlice)
				numOfSubset = unionSubset(subsets)
			}

		// Typesets are described from the union set,
		// and it is assumed that there is no union set under the product set.
		case *types.Basic:
			typName := n.String()
			if !slices.Contains(typeSlice, typName) {
				typeSlice = append(typeSlice, typName)
				subsets = getSubsetCombination(typeSlice)
			}
			numOfSubset = intersectionSubset(subsets, numOfSubset, typName)

		// TODO: support defined type
		case *types.Named:
			typName := n.String()[len("command-line-arguments."):]
			if !slices.Contains(typeSlice, typName) {
				typeSlice = append(typeSlice, typName)
				subsets = getSubsetCombination(typeSlice)
			}

		// TODO: support types and sets that are not yet supported.
		default:
			fmt.Printf("%[1]T %[1]v\n", n)
		}
	}
	printMap(numOfSubset, subsets, title)
}

func printMap(m map[string]int, subsets [][]string, title string) {
	fmt.Printf("TypeSet: %q\n", title)

	// TODO: modify empty set condition
	if len(m) == 0 {
		fmt.Printf("Empty set.\n\n")
		return
	}

	slices.SortFunc(subsets, func(a, b []string) bool { return len(a) < len(b) })
	fmt.Printf("%s\n", strings.Repeat("-", 20+len(title)))
	for _, subset := range subsets {
		k := strings.Join(subset, " ∩ ")
		fmt.Printf("%s: %d\n", k, m[k])
	}

	fmt.Println()
}
