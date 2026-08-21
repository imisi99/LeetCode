package main

import (
	"fmt"
)

func findKthSmallest(coins []int, k int) int64 {
	array := make([]int, len(coins))

	for idx := range array {
		array[idx] = 1
	}

	idx := 1
	for idx <= k {
		minIdx := 0
		minVal := coins[0] * array[0]
		for j := range array {
			val := array[j] * coins[j]
			if val < minVal {
				minVal = val
				minIdx = j
			}
		}
		idx++
		array[minIdx]++
	}

	result := array[0] * coins[0]
	for idx := range array {
		val := array[idx] * coins[idx]
		if val < result {
			result = val
		}
	}

	return int64(result)
}

func main() {
	fmt.Println(findKthSmallest([]int{3, 6, 9}, 3))
	fmt.Println(findKthSmallest([]int{5, 2}, 7))
}
