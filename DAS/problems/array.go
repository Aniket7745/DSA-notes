package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	// accesing array
	fmt.Println(arr[4])
	// modify array
	arr[2] = 6
	fmt.Println(arr[2])
	// inserting an element --O(n)
	index := 2
	value := 7

	arr = append(arr[:index], append([]int{value}, arr[index:]...)...)
	fmt.Println(arr)

	arr = append(arr[:4], append([]int{10}, arr[4:]...)...)
	fmt.Println(arr)

	// deleting
	index = 2
	arr = append(arr[:index], arr[index+1:]...)
	fmt.Println(arr)
}
