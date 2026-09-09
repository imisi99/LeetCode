package main

import (
	"fmt"
)

// Time -> 0(1) Space -> 0(1)
func countCommas(n int64) int64 {
	return max(n-999, 0) + max(n-999999, 0) + max(n-999999999, 0) + max(n-999999999999, 0) + max(n-999999999999999, 0)
}

func main() {
	fmt.Println(countCommas(1003))
	fmt.Println(countCommas(1004590))
}
