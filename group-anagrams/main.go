package main

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		chars := []rune(str)
		slices.Sort(chars)
		sortedStr := string(chars)
		m[sortedStr] = append(m[sortedStr], str)
	}
	result := make([][]string, 0, len(m))
	for _, group := range m {
		result = append(result, group)
	}
	return result
}

func main() {
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs2 := []string{""}
	strs3 := []string{"a"}
	fmt.Printf("%v\n", groupAnagrams(strs1))
	fmt.Printf("%v\n", groupAnagrams(strs2))
	fmt.Printf("%v\n", groupAnagrams(strs3))
}
