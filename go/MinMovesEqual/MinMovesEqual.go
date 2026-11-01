package main

import (
	"fmt"
	"math"
	"slices"
)

// Time -> 0(N*N*N) Space -> 0(1)
func minMoves(nums []int) int {
	count := 0
	minVal, maxVal := minMax(nums)
	for minVal != maxVal {
		skip := true
		for i := range nums {
			if (nums[i] == maxVal) && skip {
				skip = false
				continue
			}
			nums[i]++
		}
		count++
		minVal, maxVal = minMax(nums)
	}
	return count
}

func minMax(nums []int) (int, int) {
	maxVal := math.MinInt
	minVal := math.MaxInt

	for idx := range nums {
		if nums[idx] < minVal {
			minVal = nums[idx]
		}
		if nums[idx] > maxVal {
			maxVal = nums[idx]
		}
	}
	return minVal, maxVal
}

// Time -> 0(N*N) Space -> 0(1)
func minMovesII(nums []int) int {
	slices.Sort(nums)
	count := 0

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[0]
		if diff > 0 {
			for j := range nums {
				if i == j {
					continue
				}
				nums[j] += diff
			}
			count += diff
		}
	}

	return count
}

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
	fmt.Println(minMovesII([]int{1, 2, 3}))
	fmt.Println(minMovesII([]int{1, 10000000}))
	fmt.Println(minMoves([]int{1, 1, 1}))
}
