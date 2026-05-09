package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) encode(strs []string) string {
	var encoded strings.Builder
	for _, str := range strs {
		strLength := len(str)
		encoded.WriteString(fmt.Sprintf("%d#", strLength) + str)
	}
	return encoded.String()
}

func (s *Solution) decode(encoded string) []string {
	strs := []string{}
	for i := 0; i < len(encoded); {
		// Find the position of the '#' character
		j := i
		for encoded[j] != '#' {
			j++
		}
		// Extract the length of the string
		lengthStr := encoded[i:j]
		length, _ := strconv.Atoi(lengthStr)

		// Extract the string based on the length
		start := j + 1
		end := start + length
		strs = append(strs, encoded[start:end])

		i = end
	}
	return strs
}

func main() {
	strs1 := []string{"Hello", "World"}
	strs2 := []string{""}

	solution := &Solution{}
	encoded1 := solution.encode(strs1)
	encoded2 := solution.encode(strs2)

	fmt.Printf("Encoded: %s\n", encoded1)
	fmt.Printf("Encoded: %s\n", encoded2)

	decoded1 := solution.decode(encoded1)
	decoded2 := solution.decode(encoded2)

	fmt.Printf("Decoded: %v\n", decoded1)
	fmt.Printf("Decoded: %v\n", decoded2)
}
