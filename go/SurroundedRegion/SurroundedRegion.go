package main

import "fmt"

func solve(board [][]byte) {
	edits := make([][]int, 0)
	for i := range board {
		for j := range board[0] {
			if board[i][j] == 'X' {
				continue
			} else {
				explore(i, j)
			}
		}
	}
}

func explore(i, j, n, m int, board [][]byte) {
	if board[i][j] == 
}

func edge(i, j, n, m int) bool {
	if i == 0 {
		return true
	} else if i == n {
		return true
	} else if j == 0 {
		return true
	} else if j == m {
		return true
	}
	return false
}

func main() {
	board := [][]byte{
		{},
	}
	fmt.Println(board)
	solve(board)
	fmt.Println(board)
}
