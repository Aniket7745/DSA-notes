package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Max profit:", maxProfit(prices))
}

func maxProfit(prices []int) int {
	// l := len(prices)

	// brote force
	// mp := 0
	// for i := 0; i < l; i++ {
	// 	for j := i + 1; j < l; j++ {
	// 		cp := prices[j] - prices[i]
	// 		if cp > mp {
	// 			mp = cp
	// 		}
	// 	}
	// }

	mp := 0
	minP := math.MaxInt32

	for _, p := range prices {
		if p < minP {
			minP = p
		} else if (p - minP) > mp {
			mp = p - minP
		}
	}

	return mp
}
