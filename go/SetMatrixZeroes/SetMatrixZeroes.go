package main

import "fmt"

func setZeros(matrix [][]int) {
	zeros := make([][]int, 0)

	m := 0
	for m < len(matrix) {
		n := 0
		for n < len(matrix[0]) {
			if matrix[m][n] == 0 {
				zeros = append(zeros, []int{m, n})
			}
			n++
		}
		m++
	}

	for _, zero := range zeros {
		rowCol(matrix, zero[0], zero[1])
	}
}

func rowCol(matrix [][]int, m, n int) {
	start, stop := 0, len(matrix)
	for start < stop {
		matrix[start][n] = 0
		start++
	}
	start, stop = 0, len(matrix[0])
	for start < stop {
		matrix[m][start] = 0
		start++
	}
}

// Space -> 0(1)
func setZerosI(matrix [][]int) {
	zeroRow := false
	zeroCol := false

	for m := 0; m < len(matrix[0]); m++ {
		if matrix[0][m] == 0 {
			zeroRow = true
			break
		}
	}

	for n := 0; n < len(matrix); n++ {
		if matrix[n][0] == 0 {
			zeroCol = true
			break
		}
	}

	for m := 1; m < len(matrix); m++ {
		for n := 1; n < len(matrix[0]); n++ {
			if matrix[m][n] == 0 {
				matrix[m][0] = 0
				matrix[0][n] = 0
			}
		}
	}

	for m := 1; m < len(matrix); m++ {
		if matrix[m][0] == 0 {
			for n := 1; n < len(matrix[0]); n++ {
				matrix[m][n] = 0
			}
		}
	}

	for n := 1; n < len(matrix[0]); n++ {
		if matrix[0][n] == 0 {
			for m := 1; m < len(matrix); m++ {
				matrix[m][n] = 0
			}
		}
	}

	if zeroRow {
		for m := 0; m < len(matrix[0]); m++ {
			matrix[0][m] = 0
		}
	}

	if zeroCol {
		for n := 0; n < len(matrix); n++ {
			matrix[n][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}

	matrixI := [][]int{
		{0, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
	}

	setZeros(matrix)
	setZerosI(matrixI)
	fmt.Println(matrix)
	fmt.Println(matrixI)
}
