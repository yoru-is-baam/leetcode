package main

import "fmt"

func longestConsecutive(nums []int) int {
	longest := 0
	numSet := make(map[int]bool)

	// Create a set to store the unique numbers
	for _, num := range nums {
		numSet[num] = true
	}

	// Only start counting if the number is the start of a sequence
	for _, num := range nums {
		if !numSet[num-1] {
			current := num
			length := 1

			for numSet[current+1] {
				current++
				length++
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}

func main() {
	nums1 := []int{100, 4, 200, 1, 3, 2}
	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	nums3 := []int{1, 0, 2, 1}

	fmt.Println(longestConsecutive(nums1))
	fmt.Println(longestConsecutive(nums2))
	fmt.Println(longestConsecutive(nums3))
}
