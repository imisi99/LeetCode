package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func contains(nums []int) bool {
	duplicate := make(map[int]bool, 0)
	for _, num := range nums {
		if _, ok := duplicate[num]; ok {
			return true
		}
		duplicate[num] = true
	}
	return false
}

func main() {
	fmt.Println(contains([]int{0, 2, 3, 4, 4}))
	fmt.Println(contains([]int{0, 2, 3, 4, 5}))
}
