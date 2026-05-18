package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func setBit(num int) int {
	count := 0
	for num != 0 {
		bit := num % 2
		if bit == 1 {
			count++
		}
		num = num / 2
	}
	return count
}

func main() {
	fmt.Println(setBit(7))
}
