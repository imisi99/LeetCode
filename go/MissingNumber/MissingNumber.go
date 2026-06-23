package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func missingNumber(nums []int) int {
	expectedSum := 0
	actualSum := 0
	for i := 0; i <= len(nums); i++ {
		expectedSum += i
		if i != len(nums) {
			actualSum += nums[i]
		}
	}

	return expectedSum - actualSum
}

func main() {
	fmt.Println(missingNumber([]int{2, 1}))
	fmt.Println(missingNumber([]int{2, 0, 3}))
}
