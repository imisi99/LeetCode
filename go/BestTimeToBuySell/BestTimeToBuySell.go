package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func bestTime(stocks []int) int {
	profit := 0
	n := len(stocks) - 1
	bestDay := stocks[n]

	for n >= 0 {
		if bestDay-stocks[n] > profit {
			profit = bestDay - stocks[n]
		}
		if stocks[n] > bestDay {
			bestDay = stocks[n]
		}
		n--
	}
	return profit
}

func main() {
	fmt.Println(bestTime([]int{7, 1, 5, 3, 6, 4}))
}
