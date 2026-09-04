package main

import "fmt"

// Time -> 1(N) Space -> 0(N)
func firstStableIndex(nums []int, k int) int {
	lookup := make(map[int][]int, 0)

	maxVal := nums[0]
	for idx := range nums {
		if nums[idx] > maxVal {
			maxVal = nums[idx]
		}
		lookup[idx] = []int{maxVal}
	}

	idx := len(nums) - 1
	minVal := nums[idx]
	for idx >= 0 {
		if nums[idx] < minVal {
			minVal = nums[idx]
		}
		lookup[idx] = append(lookup[idx], minVal)
		idx--
	}

	for idx := range nums {
		score := lookup[idx]
		instable := score[0] - score[1]
		if instable <= k {
			return idx
		}
	}
	return -1
}

func main() {
	fmt.Println(firstStableIndex([]int{5, 0, 1, 4}, 3))
	fmt.Println(firstStableIndex([]int{3, 2, 1}, 1))
	fmt.Println(firstStableIndex([]int{0}, 0))
	fmt.Println(firstStableIndex([]int{6, 1, 4}, 5))
}
