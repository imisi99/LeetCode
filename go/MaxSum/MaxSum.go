package main

import (
	"fmt"
)

// Time => 0(N) Space => 0(1)
func maxSum(nums []int, k int) int {
	var coins [31]int
	const c = 1_000_000_007

	for _, num := range nums {
		for b := range len(coins) {
			coins[b] += ((num >> b) & 1)
		}
	}

	total := 0
	for ; k > 0; k-- {
		curr := 0
		for b := range len(coins) {
			if coins[b] > 0 {
				curr |= (1 << b)
				coins[b]--
			}
		}

		total = (total + (curr * curr)) % c
	}
	return total
}

func main() {
	fmt.Println(maxSum([]int{2, 6, 5, 8}, 2))
	fmt.Println(maxSum([]int{4, 5, 4, 7}, 3))
}
