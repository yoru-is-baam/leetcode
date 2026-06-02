package main

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}

	// Sort the input array to make it easier to avoid duplicates
	slices.Sort(nums)

	// Iterate through the sorted array
	for i := range nums {
		// Skip duplicate elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a, left, right := nums[i], i+1, len(nums)-1
		for left < right {
			if nums[left]+nums[right] == -a {
				result = append(result, []int{a, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if nums[left]+nums[right] < -a {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	nums2 := []int{0, 1, 1}
	nums3 := []int{0, 0, 0}

	fmt.Println("Example 1:", threeSum(nums1))
	fmt.Println("Example 2:", threeSum(nums2))
	fmt.Println("Example 3:", threeSum(nums3))
}
