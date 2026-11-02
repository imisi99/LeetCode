package main

import (
	"fmt"
	"slices"
)

// Time -> 0(NlogN) Space -> 0(1)
func minMoves(nums []int) int {
	slices.Sort(nums)
	mid := 0
	if len(nums)%2 == 1 {
		mid = nums[len(nums)/2]
	} else {
		mid = (nums[len(nums)/2] + nums[((len(nums)/2)-1)]) / 2
	}

	count := 0
	for i := range nums {
		val := nums[i] - mid
		if val < 0 {
			val *= -1
		}
		count += val
	}
	return count
}

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
}
