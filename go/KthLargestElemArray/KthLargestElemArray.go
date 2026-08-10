package main

import "fmt"

// Time -> 0(N) Space -> (N)  N is the max num in the array
func largest(array []int, k int) int {
	max := 0
	for idx := range array {
		if array[idx] > max {
			max = array[idx]
		}
	}

	prefill := make([]int, max+1)

	for _, val := range array {
		prefill[val]++
	}

	fmt.Println(prefill)

	idx := max
	for k > 1 {
		if prefill[idx] >= k {
			break
		}
		k -= prefill[idx]
		idx--
	}

	return idx
}

func main() {
	fmt.Println(largest([]int{2, 3, 4, 6, 6, 9}, 3))
}
