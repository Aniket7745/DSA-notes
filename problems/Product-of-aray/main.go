package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4}
	result1 := productExceptSelf(nums1)
	fmt.Println("Output for [1, 2, 3, 4]:", result1) // Expected: [24 12 8 6]

	nums2 := []int{-1, 1, 0, -3, 3}
	result2 := productExceptSelf(nums2)
	fmt.Println("Output for [-1, 1, 0, -3, 3]:", result2) // Expected: [0 0 9 0 0]
}

func productExceptSelf(num []int) int {
	return 0
}
