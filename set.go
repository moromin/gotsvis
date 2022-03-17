package gotsvis

import (
	"fmt"
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

func existType(numOfSubset map[string]int, typName string, tilde bool) bool {
	var ok1, ok2 bool
	if tilde {
		_, ok1 = numOfSubset[typName]
		_, ok2 = numOfSubset[typName[1:]]
	} else {
		_, ok1 = numOfSubset[typName]
		_, ok2 = numOfSubset[fmt.Sprintf("%s%s", "~", typName)]
	}

	if !ok1 && !ok2 {
		return false
	}
	return true
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
	// check empty set
	tilde := strings.HasPrefix(typName, "~")
	if !existType(numOfSubset, typName, tilde) {
		return map[string]int{}
	}

	res := numOfSubset
	for _, subset := range subsets {
		n := len(subset)
		if n == 1 && subset[0] == typName {
			if tilde {
				res[typName] = 100
			} else {
				res[typName] = 10
			}
		} else if n == 2 {
			if tilde && slices.Contains(subset, typName) && slices.Contains(subset, typName[1:]) ||
				!tilde && slices.Contains(subset, typName) && slices.Contains(subset, fmt.Sprintf("%s%s", "~", typName)) {
				res[strings.Join(subset, " ∩ ")] = 10
			}
		}
	}
	return res
}
