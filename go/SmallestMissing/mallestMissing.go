package main

import (
	"fmt"
	"slices"
)

// Time -> 0(N) Space -> 0(N)
func missingMultiple(nums []int, k int) int {
	lookup := make(map[int]bool, 0)

	for idx := range nums {
		lookup[nums[idx]] = true
	}

	val := 1
	for val <= len(nums) {
		if _, exists := lookup[val*k]; !exists {
			return val * k
		}
		val++
	}

	return k * (len(nums) + 1)
}

// Time -> 0(NlogN) Space -> 0(1)
func missingMultipleI(nums []int, k int) int {
	slices.Sort(nums)

	mul := k
	for _, num := range nums {
		if num > mul {
			return mul
		} else if num == mul {
			mul += k
		}
	}

	return mul
}

func main() {
	fmt.Println(missingMultiple([]int{2, 4, 6}, 2))
	fmt.Println(missingMultipleI([]int{6, 4, 2}, 2))
}
