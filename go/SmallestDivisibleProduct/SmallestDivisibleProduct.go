package main

import (
	"fmt"
	"strconv"
)

// Time -> 0(1) Space -> 0(1)
func smallest(n, t int) int {
	if n%10 == 0 {
		return n
	}

	k := (n + 10) % 10
	end := k * 10

	for n < end {
		val := strconv.Itoa(n)
		prod := 1

		for v := range val {
			intv, _ := strconv.Atoi(string(val[v]))
			prod *= intv
		}

		if prod%t == 0 {
			return n
		}
		n++
	}

	return end
}

func main() {
	fmt.Println(smallest(10, 3))
	fmt.Println(smallest(14, 3))
}
