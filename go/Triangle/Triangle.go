package main

import "fmt"

// Time -> 0(N^2) Space -> 0(N^2)
func minimumTotal(triangle [][]int) int {
	memo := make(map[[2]int]int, 0)
	var recurse func(row, col int) int
	recurse = func(row, col int) int {
		if row >= len(triangle) || col >= len(triangle[row]) {
			return 0
		}

		if val, exists := memo[[2]int{row, col}]; exists {
			return val
		}

		samecol := recurse(row+1, col)
		diffcol := recurse(row+1, col+1)

		memo[[2]int{row, col}] = triangle[row][col] + min(samecol, diffcol)
		return memo[[2]int{row, col}]
	}

	return recurse(0, 0)
}

func main() {
	fmt.Println(minimumTotal([][]int{
		{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3},
	}))
}
