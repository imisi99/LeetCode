package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func isGood(nums []int) bool {
	base := len(nums) - 1
	baseMap := make([]int, base+1)
	for _, num := range nums {
		if num <= 0 || num > base {
			return false
		}

		baseMap[num]++

		if baseMap[num] > 2 || (baseMap[num] == 2 && num != base) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isGood([]int{1, 2, 3, 3}))
	fmt.Println(isGood([]int{0, 1, 2, 3}))
}
