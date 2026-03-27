package main

import (
	"fmt"
)

func climbStairs(n int) int {
	memo := make(map[int]int)
	return countStairs(0, n, memo)
}

// Time -> 0(N) Space -> 0(N)
func countStairs(idx, n int, memo map[int]int) int {
	if idx >= n {
		if idx == n {
			return 1
		}
		return 0
	}

	if val, ok := memo[idx]; ok {
		return val
	}

	oneStep := countStairs(idx+1, n, memo)
	twoStep := countStairs(idx+2, n, memo)

	memo[idx] = oneStep + twoStep

	return memo[idx]
}

func main() {
	fmt.Println(climbStairs(44))
}
