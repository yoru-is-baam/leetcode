package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}

func main() {
	// nums := []int{2, 7, 11, 15}
	// target := 9
	nums := []int{3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}
