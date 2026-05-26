package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func majorityElement(nums []int) int {
	counter := make(map[int]int, 0)

	for _, num := range nums {
		counter[num]++
		if counter[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 2}))
}
