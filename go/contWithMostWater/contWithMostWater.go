package main

import "fmt"

// Time -> 0(N^2) Space -> 0(N)
func contWithMostWater(array []int, idx int) int {
	if idx == len(array)-2 {
		area := min(array[len(array)-1], array[len(array)-2])
		return area
	}

	area := findMax(array, idx)
	return max(area, contWithMostWater(array, idx+1))
}

func findMax(array []int, idx int) int {
	max := 0
	j := idx

	for i := idx; i < len(array); i++ {
		width := i - j
		height := min(array[j], array[i])
		area := width * height
		if area > max {
			max = area
		}
	}
	return max
}

// Time -> 0(N) Space -> 0(1)
func soln2(array []int) int {
	area := 0
	first, end := 0, len(array)-1

	for first < end {
		currArea := min(array[first], array[end]) * (end - first)
		area = max(area, currArea)

		if array[first] < array[end] {
			first++
		} else {
			end--
		}
	}
	return area
}

func main() {
	fmt.Println(contWithMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 0))
	fmt.Println(soln2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
