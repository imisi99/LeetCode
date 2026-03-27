package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func removeDuplicates(nums []int) int {
	idxToswap := 0
	idx := 0

	for idx < len(nums) {
		currIdx := idx
		for idx < len(nums) && nums[currIdx] == nums[idx] {
			idx++
		}

		if idx-currIdx-1 >= 1 {
			nums[idxToswap] = nums[currIdx]
			nums[idxToswap+1] = nums[currIdx]
			idxToswap += 2
		} else {
			nums[idxToswap] = nums[currIdx]
			idxToswap++
		}
	}

	return idxToswap
}

func main() {
	array := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(array)
	fmt.Println(removeDuplicates(array))
	fmt.Println(array)
}
