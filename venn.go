package gotsvis

import (
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/exp/slices"
)

func contains(s []string, v string) bool {
	for _, elem := range s {
		if elem == v {
			return true
		}
	}
	return false
}

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

func calcNumOfSubset(subsets [][]string) map[string]int {
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

func Venn(title string, s []types.Type) {
	// Name:size
	// res := make(map[string]int)

	// set := map[string][]string{
	// 	"Union":        []string{},
	// 	"Intersection": []string{},
	// }

	typeSlice := []string{}
	subsets := [][]string{}

	for _, elem := range s {
		switch n := elem.(type) {
		case *types.Union:
			for i := 0; i < n.Len(); i++ {
				typeSlice = append(typeSlice, n.Term(i).String())
			}
			subsets = getSubsetCombination(typeSlice)
			numOfSubset := calcNumOfSubset(subsets)
			printMap(numOfSubset, subsets, title)
		}
	}
}

func printMap(m map[string]int, subsets [][]string, title string) {
	slices.SortFunc(subsets, func(a, b []string) bool { return len(a) < len(b) })

	fmt.Printf("TypeSet: %q\n", title)
	fmt.Printf("%s\n", strings.Repeat("-", 20+len(title)))
	for _, subset := range subsets {
		k := strings.Join(subset, " ∩ ")
		fmt.Printf("%s: %d\n", k, m[k])
	}

	fmt.Println()
}
