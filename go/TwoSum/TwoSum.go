package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func twoSums(nums []int, target int) []int {
	set := make(map[int]int, 0)
	for i, val := range nums {
		set[val] = i
	}

	for i, val := range nums {
		diff := target - val
		if val, ok := set[diff]; ok && i != val {
			return []int{val, i}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSums([]int{2, 7, 11, 15}, 9))
}
