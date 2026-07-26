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

func explore(i, j int) {
}

func main() {
	board := [][]byte{
		{},
	}
	fmt.Println(board)
	solve(board)
	fmt.Println(board)
}
