package main

import (
	"fmt"
	"math"
)

// Time -> 0(N) Space -> 0(N)
func reverse(num int) int {
	array := make([]int, 0)
	for num != 0 {
		bit := num % 2
		array = append(array, bit)
		num = num / 2
	}

	for len(array) < 32 {
		array = append(array, 0)
	}

	power := 0
	idx := len(array) - 1
	val := 0
	for idx >= 0 {
		pow := math.Pow(float64(2), float64(power))
		val += (int(pow) * array[idx])
		power++
		idx--
	}

	return val
}

func main() {
	fmt.Println(reverse(2147483644))
}
