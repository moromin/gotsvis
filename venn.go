package gotsvis

import (
	"fmt"
	"go/types"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func getVennDiagram(m map[string]int) {
	baseURL := "http://chart.apis.google.com/chart"
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()

	names := []string{}
	numbers := []string{}
	for k, v := range m {
		names = append(names, k)
		numbers = append(numbers, strconv.Itoa(v))
	}

	q.Add("cht", "v")
	q.Add("chs", "300x300")
	q.Add("chd", fmt.Sprintf("t:%s", strings.Join(numbers, ",")))
	q.Add("chdl", strings.Join(names, "|"))
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("venn.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(out, img); err != nil {
		log.Fatal(err)
	}
}

func calcSet(s [][]string) map[string]int {
	res := make(map[string]int)
	for _, typs := range s {
		n := len(typs)
		if n == 1 {
			if strings.HasPrefix(typs[0], "~") {
				res[typs[0]] = 100
			} else {
				res[typs[0]] = 10
			}
		} else {
			res[strings.Join(typs, "∩")] = 0
		}
	}
	return res
}

func Venn(title string, s []types.Type) {
	// set := map[string][]string{
	// 	"Union":        []string{},
	// 	"Intersection": []string{},
	// }

	// Name:size
	// res := make(map[string]int)
	typeSlice := make([]string, 0)

	for _, elem := range s {
		switch n := elem.(type) {
		case *types.Union:
			for i := 0; i < n.Len(); i++ {
				term := n.Term(i)
				typeName := term.String()
				typeSlice = append(typeSlice, typeName)
				// if term.Tilde() {
				// 	res[typeName] = 100
				// } else {
				// 	if term.String() != term.Type
				// 	if _, ok := res[typeName]; !ok {

				// 	}
				// }
			}
			// getSubsetCombination(typeSlice)
			subset := getSubsetCombination(typeSlice)
			res := calcSet(subset)
			getVennDiagram(res)
		}
	}
}
