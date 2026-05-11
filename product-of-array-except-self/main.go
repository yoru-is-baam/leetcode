package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	cumulativeProduct := 1

	for i := range nums {
		result[i] = cumulativeProduct
		cumulativeProduct *= nums[i]
	}

	cumulativeProduct = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= cumulativeProduct
		cumulativeProduct *= nums[i]
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{-1, 1, 0, -3, 3}

	fmt.Println(productExceptSelf(nums1)) // Output: [24, 12, 8, 6]
	fmt.Println(productExceptSelf(nums2)) // Output: [0, 0, 9, 0, 0]
}
