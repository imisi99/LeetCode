package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func moveZeros(nums []int) {
	idxToSwap := 0
	idx := 0
	for idx < len(nums) {
		if nums[idx] != 0 {
			nums[idxToSwap], nums[idx] = nums[idx], nums[idxToSwap]
			idxToSwap++
		}
		idx++
	}
}

func main() {
	array := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	fmt.Println(array)
	moveZeros(array)
	fmt.Println(array)
}
