package main

import (
	"fmt"
	"strconv"
)

// Time -> 0(N^2) Space -> 0(N)
func addDigits(num int) int {
	for len(fmt.Sprintf("%v", num)) != 1 {
		runSum := 0
		val := fmt.Sprintf("%v", num)
		for idx := range val {
			add, _ := strconv.Atoi(string(val[idx]))
			runSum += add
		}
		num = runSum
	}
	return num
}

func main() {
	fmt.Println(addDigits(121))
}
