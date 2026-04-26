package main

import "fmt"

func spiralMatrix(matrix [][]int) []int {
	array := make([]int, 0)
	spiral(0, len(matrix[0])-1, 0, len(matrix)-1, matrix, &array)
	return array
}

// Time -> 0(M*N) Space ->  0(M*N)
func spiral(startCol, endCol, startRow, endRow int, matrix [][]int, array *[]int) {
	for startRow <= endRow && startCol <= endCol {
		// go right
		right := startCol
		for right <= endCol {
			*array = append(*array, matrix[startRow][right])
			right++
		}

		startRow++
		if startRow > endRow {
			break
		}

		// go down
		down := startRow
		for down <= endRow {
			*array = append(*array, matrix[down][endCol])
			down++
		}

		endCol--
		if endCol < startCol {
			break
		}

		// go left
		left := endCol
		for left >= startCol {
			*array = append(*array, matrix[endRow][left])
			left--
		}

		endRow--
		if startRow > endRow {
			break
		}

		// go up
		up := endRow
		for up >= startRow {
			*array = append(*array, matrix[up][startCol])
			up--
		}
		startCol++
	}
}

func main() {
	fmt.Println(spiralMatrix([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
		{21, 22, 23, 24},
	}))
}
