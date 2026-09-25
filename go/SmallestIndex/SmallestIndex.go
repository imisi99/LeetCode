package main

import "fmt"

// Time -> 1(Nd) Space -> 0(d)
func smallestIndex(nums []int) int {
	for idx := range nums {
		val := fmt.Sprint(nums[idx])
		sum := 0
		for i := range val {
			sum += int(val[i] - '0')
		}
		if sum == idx {
			return idx
		}
	}

	return -1
}

func main() {
	fmt.Println(smallestIndex([]int{1, 3, 2}))
	fmt.Println(smallestIndex([]int{1, 10, 11}))
	fmt.Println(smallestIndex([]int{1, 2, 3}))
}
