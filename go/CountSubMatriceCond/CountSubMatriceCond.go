package main

import "fmt"

// Time -> 0(max(N, M)) Space -> 0(1)
func numberOfSubmatrices(grid [][]byte) int {
	var count int
	countX := 0
	countY := 0

	bottom, right := 0, 0
	for bottom < len(grid) || right < len(grid[0]) {
		if right == len(grid[0]) {
			right = len(grid[0]) - 1
		}

		if bottom == len(grid) {
			bottom = len(grid) - 1
		}

		for i := 0; i <= bottom; i++ {
			count += check(i, right, &countX, &countY, grid)
		}
		bottom++
		right++
	}

	return count
}

func check(i, j int, countX, countY *int, grid [][]byte) int {
	val := grid[i][j]
	if val == 'X' {
		*countX++
	}

	if val == 'Y' {
		*countY++
	}

	if *countX == *countY && *countX != 0 {
		return 1
	}
	return 0
}

func main() {
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}}))
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'X'}, {'X', 'Y'}}))
	fmt.Println(numberOfSubmatrices([][]byte{{'.', '.'}, {'Y', 'X'}}))
}
