package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	testCase1 := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	results := findAnagrams(testCase1)
	fmt.Println(results)
}

func findAnagrams(s []string) map[string][]string {
	sortedSym := func(st string) string {
		rs := []rune(st)
		sort.Slice(rs, func(i, j int) bool {return rs[i] < rs[j]})
		return string(rs)
	}

	type void struct{}
	groups := make(map[string]map[string]void)
	firstSeen := make(map[string]string)

	for i := 0; i < len(s); i++ {
		w := strings.ToLower(string(s[i]))
		if w == "" {
			continue
		}

		key := sortedSym(w)
		if _, ok := groups[key]; !ok {
			groups[key] = make(map[string]void, 1)
			firstSeen[key] = w
		}
		groups[key][w] = void{}
	}

	result := make(map[string][]string)
	for sig, set := range groups {
		list := make([]string, 0, len(set))
		for w := range set {
			list = append(list, w)
		}
		if len(list) < 2 {
			continue
		}
		sort.Strings(list)
		result[firstSeen[sig]] = list
	}
	return result
}
