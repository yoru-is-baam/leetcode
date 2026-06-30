package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	target2 := 2

	fmt.Println(search(nums1, target1)) // Output: 4
	fmt.Println(search(nums1, target2)) // Output: -1
}
