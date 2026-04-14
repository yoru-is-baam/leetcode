package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Printf("Array 1: %v\n", containsDuplicate(nums1))
	fmt.Printf("Array 2: %v\n", containsDuplicate(nums2))
	fmt.Printf("Array 3: %v\n", containsDuplicate(nums3))
}
