package main

import (
	"fmt"
	"math"
	"strconv"
)

// Time -> 0(N) Space -> 0(N)
func reverse(x int) int {
	val := fmt.Sprintf("%d", x)

	runes := []rune(val)
	start := 0
	if runes[0] == rune('-') {
		start = 1
	}

	for i, j := start, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	val = string(runes)
	valInt, _ := strconv.ParseInt(val, 10, 64)

	if valInt < int64(math.Pow(float64(-2), float64(31))) || valInt > int64(math.Pow(float64(2), float64(31)))-1 {
		return 0
	}

	return int(valInt)
}

func main() {
	fmt.Println(reverse(-456))
}
