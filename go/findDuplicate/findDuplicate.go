package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func findDuplicate(nums []int) int {
	actualSum := 0
	expSum := 0

	i := 1
	for i < len(nums) {
		expSum += i
		i++
	}

	for idx := range nums {
		actualSum += nums[idx]
	}

	return actualSum - expSum
}

func main() {
	fmt.Println(findDuplicate([]int{1, 2, 3, 4, 5, 6, 6}))
}
