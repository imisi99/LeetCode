package main

import "fmt"

// Time => 0(N) Space => 0(1)
func sumOfSquares(nums []int) int {
	sumSquares := 0
	n := len(nums)
	for i := range nums {
		if n%(i+1) == 0 {
			sumSquares += (nums[i] * nums[i])
		}
	}
	return sumSquares
}

func main() {
	fmt.Println(sumOfSquares([]int{1, 2, 3, 4}))
	fmt.Println(sumOfSquares([]int{2, 7, 1, 19, 18, 3}))
}
