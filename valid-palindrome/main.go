package main

import (
	"fmt"
)

func isAlphaNumeric(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func isPalindrome(s string) bool {
	lP, rP := 0, len(s)-1
	for lP < rP {
		if !isAlphaNumeric(s[lP]) {
			lP++
			continue
		}
		if !isAlphaNumeric(s[rP]) {
			rP--
			continue
		}
		if toLower(s[lP]) != toLower(s[rP]) {
			return false
		}
		lP++
		rP--
	}
	return true
}

func main() {
	s1 := "A man, a plan, a canal: Panama"
	s2 := "race a car"
	s3 := " "
	fmt.Printf("%v\n", isPalindrome(s1))
	fmt.Printf("%v\n", isPalindrome(s2))
	fmt.Printf("%v\n", isPalindrome(s3))
}
