package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println("Result:", result)
}

// brootforce O(n^2)
// func twoSum(nums []int, target int) []int {
// 	l := len(nums)
//
// 	for i := 0; i < l; i++ {
// 		for j := i + 1; j < l; j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }

// map function O(n)
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		if j, x := mp[diff]; x {
			return []int{j, i}
		}

		mp[nums[i]] = i
	}
	return []int{}
}
