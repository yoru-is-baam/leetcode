package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, char := range s {
		count[char]++
	}

	for _, char := range t {
		count[char]--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Printf("Is '%s' an anagram of '%s'? %v\n", s, t, isAnagram(s, t))

	s2 := "rat"
	t2 := "car"
	fmt.Printf("Is '%s' an anagram of '%s'? %v\n", s2, t2, isAnagram(s2, t2))
}
