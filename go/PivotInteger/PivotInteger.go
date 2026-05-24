package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func pivotInteger(n int) int {
	totSum := 0
	start := 1

	for start <= n {
		totSum += start
		start++
	}

	start = 1
	runSum := 0
	for start <= n {
		runSum += start
		if runSum == totSum {
			return start
		}
		totSum -= start
		start++
	}

	return -1
}

func main() {
	fmt.Println(pivotInteger(4))
}
