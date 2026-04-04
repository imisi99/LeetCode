package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func sortColors(nums []int) {
	red, blue := 0, 2

	blueIdx := len(nums) - 1
	redIdx := 0
	idx := 0
	for idx <= blueIdx {
		val := nums[idx]
		switch val {
		case red:
			nums[idx], nums[redIdx] = nums[redIdx], nums[idx]
			redIdx++
			idx++
		case blue:
			nums[idx], nums[blueIdx] = nums[blueIdx], nums[idx]
			blueIdx--
		default:
			idx++
		}
	}
}

func main() {
	array := []int{2, 0, 2, 0, 2, 1, 1, 0}
	fmt.Println(array)
	sortColors(array)
	fmt.Println(array)
}
