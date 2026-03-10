package main

import "fmt"

// RemoveDuplicateInPlace  Time -> 0(N) Space -> 0(1)
func RemoveDuplicateInPlace(array []int) int {
	k := 1
	idxToSwap, lastval := 1, array[0]

	for i := range array {
		if array[i] != lastval {
			array[idxToSwap] = array[i]
			lastval = array[i]
			idxToSwap++
			k++
		}
	}
	return k
}

func main() {
	array := []int{1, 1, 2, 3, 3, 4, 5, 5, 8, 9, 9}
	fmt.Printf("Array -> %v", array)
	k := RemoveDuplicateInPlace(array)
	fmt.Printf("Array -> %v num of unique item -> %v", array, k)
}
