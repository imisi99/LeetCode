package main

import "fmt"

// Time -> 0(log(M * N))
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	start, end := 0, len(matrix)-1
	bestMid := 0
	for start <= end {
		mid := (start + end) / 2

		val := matrix[mid][0]
		if val == target {
			return true
		} else if val < target {
			bestMid = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	start, end = 0, len(matrix[0])-1
	for start <= end {
		mid := (start + end) / 2

		val := matrix[bestMid][mid]
		if val == target {
			return true
		} else if val < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{},
	}
	fmt.Println(searchMatrix(matrix, 8))
}
