package main

import (
	"fmt"
)

func robbery(profits []int) int {
	memo := make([]int, len(profits))
	for i := range memo {
		memo[i] = -1
	}
	return max(findMax(0, memo, profits), findMax(1, memo, profits))
}

// Time -> 0(N) Space -> 0(N)
func findMax(i int, memo, profits []int) int {
	if i >= len(profits) {
		return 0
	}

	if memo[i] != -1 {
		return memo[i]
	}

	memo[i] = profits[i] + max(findMax(i+2, memo, profits), findMax(i+3, memo, profits))
	return memo[i]
}

func main() {
	fmt.Println(robbery([]int{1, 2, 3, 1}))
}
