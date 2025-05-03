package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 1}
	fmt.Println("Contains duplicate:", containsDuplicate(nums1)) // Output: true

	nums2 := []int{1, 2, 3, 4}
	fmt.Println("Contains duplicate:", containsDuplicate(nums2)) // Output: false

	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println("Contains duplicate:", containsDuplicate(nums3))
}

func containsDuplicate(nums []int) bool {
	// broot force
	// for i := 0; i < len(nums); i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i] == nums[j] {
	// 			return true
	// 		}
	// 	}
	// }

	// shorting O(n log n)
	// sort.Ints(nums)
	//
	// for i := 1; i < len(nums); i++ {
	// 	if nums[i] == nums[i-1] {
	// 		return true
	// 	}
	// }

	// map
	mp := make(map[int]bool)
	for _, num := range nums {
		if mp[num] {
			return true
		}
		mp[num] = true
	}

	return false
}
