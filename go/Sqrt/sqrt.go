package main

import "fmt"

// Time -> 0(log(N)) Space -> 0(1)
func mySqrt(x int) int {
	start, stop := 0, x
	best := 0
	for start <= stop {
		mid := (start + stop) / 2
		midSquare := mid * mid
		if midSquare == x {
			return mid
		} else if midSquare > x {
			stop = mid - 1
		} else {
			start = mid + 1
		}
		if midSquare < x {
			best = mid
		}
	}
	return best
}

func main() {
	fmt.Println(mySqrt(15))
}
