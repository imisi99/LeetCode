package main

import "fmt"

// Time -> 0(N^2) Space -> 0(N)
func bulbSwitch(n int) int {
	array := make([]int, n)

	multiplier := 1
	for idx := range array {
		for idx < len(array) {
			inverse(array, idx)
			idx += multiplier
		}
		multiplier++
	}

	counter := 0

	for idx := range array {
		if array[idx] == 1 {
			counter++
		}
	}
	return counter
}

func inverse(array []int, idx int) {
	if array[idx] == 0 {
		array[idx] = 1
		return
	}
	array[idx] = 0
}

func main() {
	fmt.Println(bulbSwitch(8))
}
