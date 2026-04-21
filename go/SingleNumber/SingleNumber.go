package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func singleNumber(nums []int) int {
	val := 0
	for _, num := range nums {
		val ^= num
	}
	return val
}

func main() {
	fmt.Println(singleNumber([]int{1, 1, 2, 3, 4, 3, 2}))
}
