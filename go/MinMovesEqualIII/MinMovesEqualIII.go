// Time -> 0(N) Space -> 0(1)
package main

import "fmt"

func minMoves(nums []int) int {
	count := 0
	maxVal := 0

	for i := range nums {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}

	for i := range nums {
		count += (maxVal - nums[i])
	}

	return count
}

func main() {
	fmt.Println(minMoves([]int{1, 2, 3}))
	fmt.Println(minMoves([]int{4, 4, 5}))
}
