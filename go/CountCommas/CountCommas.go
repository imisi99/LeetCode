package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func countCommas(n int) int {
	num := n / 1000

	if num < 1 {
		return 0
	}

	count := ((num - 1) * 1000) + 1
	count += (n % 1000)

	return count
}

func main() {
	fmt.Println(countCommas(998))
	fmt.Println(countCommas(1003))
	fmt.Println(countCommas(4003))
}
