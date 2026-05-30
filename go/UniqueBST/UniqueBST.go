package main

import "fmt"

func numTrees(n int) int {
	sum := 0
	recurse(1, n, &sum)
	return sum
}

func recurse(idx, n int, sum *int) {
	if idx == n {
		*sum++
		return
	}
	if idx > n {
		return
	}

	recurse(idx+1, n, sum) // Go left
	recurse(idx+1, n, sum) // Go right
	recurse(idx+2, n, sum) // Construct left & right
}

func main() {
	fmt.Println(numTrees(4))
}
