package main

import "fmt"

// RemoveDuplicateInPlace  Time -> 0(N) Space -> 0(1)
func RemoveDuplicateInPlace(array []int) int {
	k := 1
	idxToSwap, idx, prev := 1, 1, 0

	for idx < len(array) {
		for array[idx] == array[prev] {
			if idx == len(array)-1 {
				break
			}
			idx++
			prev++
		}
		array[idxToSwap] = array[idx]
		idxToSwap++
		idx++
		prev++
		k++
	}

	return k
}

func main() {
	array := []int{1, 1, 2, 3, 3, 4, 4, 4, 7, 11}
	fmt.Printf("Array -> %v", array)
	k := RemoveDuplicateInPlace(array)
	fmt.Printf("Array -> %v num of unique item -> %v", array, k)
}
