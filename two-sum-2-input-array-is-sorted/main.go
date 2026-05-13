package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]+nums[right] == target {
			return []int{left + 1, right + 1}
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

func main() {
	// non-decreasing order
	nums1 := []int{2, 7, 11, 15}
	nums2 := []int{2, 3, 4}
	nums3 := []int{-1, 0}

	fmt.Printf("%v\n", twoSum(nums1, 9))
	fmt.Printf("%v\n", twoSum(nums2, 6))
	fmt.Printf("%v\n", twoSum(nums3, -1))
}
