package main

import (
	"fmt"
)

func StockPicker(arr []int) int {
	if len(arr) < 2 {
		return -1 // Not enough data for a valid solution
	}

	maxProfit := 0
	minPrice := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > minPrice {
			profit := arr[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			minPrice = arr[i]
		}
	}

	if maxProfit == 0 {
		return -1 // No profitable trade found
	}
	return maxProfit
}

func main() {
	// Output: 5
	fmt.Println(StockPicker([]int{10, 12, 4, 5, 9}))
}
