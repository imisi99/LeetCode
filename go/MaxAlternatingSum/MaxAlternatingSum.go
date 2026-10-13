package main

import (
	"fmt"
	"slices"
)

// Time => 0(NlogN) Space => 0(1)
func maxAlternatingSum(nums []int) int64 {
	for i := range nums {
		nums[i] = (nums[i] * nums[i])
	}

	slices.Sort(nums)

	k := len(nums) / 2

	subtract := nums[:k]
	add := nums[k:]

	var total int64
	for _, val := range add {
		total += int64(val)
	}

	for _, val := range subtract {
		total -= int64(val)
	}
	return total
}

func main() {
	fmt.Println(maxAlternatingSum([]int{1, 2, 3}))
	fmt.Println(maxAlternatingSum([]int{1, -1, 2, -2, 3, -3}))
}
