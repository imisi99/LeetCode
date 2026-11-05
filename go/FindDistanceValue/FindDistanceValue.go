package main

import "fmt"

// Time -> 0(MN) Space -> 0(1)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	count := 0

	for i := range arr1 {
		addVal := 1
		for j := range arr2 {
			if abs(arr1[i]-arr2[j]) <= d {
				addVal = 0
				break
			}
		}
		count += addVal
	}
	return count
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}
