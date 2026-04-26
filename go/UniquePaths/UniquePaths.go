package main

import "fmt"

func uniquePaths(m, n int) int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}
	return searchPaths(0, 0, m, n, memo)
}

// Time -> 0(M*N) Space -> 0(M*N)
func searchPaths(i, j, m, n int, memo [][]int) int {
	if i == m || j == n {
		return 0
	}
	if i == m-1 && j == n-1 {
		return 1
	}

	if memo[i][j] != 0 {
		return memo[i][j]
	}

	path := 0

	path += searchPaths(i+1, j, m, n, memo)
	path += searchPaths(i, j+1, m, n, memo)

	memo[i][j] = path

	return memo[i][j]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
