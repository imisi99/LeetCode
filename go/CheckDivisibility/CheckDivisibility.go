package main

import (
	"fmt"
)

// Time -> 0(N) Space -> 0(1)
func checkDivisibility(n int) bool {
	val := fmt.Sprint(n)
	prod, sum := 1, 0

	for idx := range val {
		num := int(val[idx] - '0')
		prod *= num
		sum += num
	}

	return n%(prod+sum) == 0
}

func main() {
	fmt.Println(checkDivisibility(99))
	fmt.Println(checkDivisibility(23))
}
