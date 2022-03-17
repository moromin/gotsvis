package gotsvis

import (
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/exp/slices"
)

func getSubsetCombination(s []string) [][]string {
	n := len(s)
	res := make([][]string, 0)

	if n == 0 {
		return res
	}

	for i := 1; i < (1 << n); i++ {
		set := make([]string, 0)
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				set = append(set, s[j])
			}
		}
		res = append(res, set)
	}
	return res
}

func unionSubset(subsets [][]string) map[string]int {
	res := make(map[string]int)

	for _, subset := range subsets {
		if len(subset) == 1 {
			if strings.HasPrefix(subset[0], "~") {
				res[subset[0]] = 100
			} else {
				res[subset[0]] = 10
			}
		} else {
			res[strings.Join(subset, " ∩ ")] = 0
		}
	}
	return res
}

func intersectionSubset(subsets [][]string, numOfSubset map[string]int, typName string) map[string]int {
	if _, ok := numOfSubset[typName]; !ok {
		return map[string]int{}
	}

	res := numOfSubset
	for _, subset := range subsets {
		if !(len(subset) == 1 && subset[0] == typName) {
			res[strings.Join(subset, " ∩ ")] = 0
		}
	}
	return res
}

func Venn(title string, s []types.Type, du map[string]string) {
	typeSlice := []string{}
	subsets := [][]string{}
	numOfSubset := make(map[string]int)

	for _, elem := range s {
		switch n := elem.(type) {
		case *types.Union:
			// TODO: ~*** intersection other, (ex. int | string; ~int)
			// if n.Len() == 1 {}

			for i := 0; i < n.Len(); i++ {
				typeSlice = append(typeSlice, n.Term(i).String())
			}
			subsets = getSubsetCombination(typeSlice)
			numOfSubset = unionSubset(subsets)

		// Typesets are described from the union set,
		// and it is assumed that there is no union set under the product set.
		case *types.Basic:
			typName := n.String()
			if !slices.Contains(typeSlice, typName) {
				typeSlice = append(typeSlice, typName)
				subsets = getSubsetCombination(typeSlice)
			}
			numOfSubset = intersectionSubset(subsets, numOfSubset, typName)

		// TODO: support types and sets that are not yet supported.
		default:
			fmt.Printf("%[1]T %[1]v\n", n)
		}
		// fmt.Println(subsets)
		// fmt.Println(numOfSubset)
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
