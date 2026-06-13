package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func canWinNim(n int) bool {
	memo := make(map[int]bool, 0)

	var playOptimal func(n int) bool
	playOptimal = func(n int) bool {
		if n <= 0 {
			return false
		}

		if val, exists := memo[n]; exists {
			return val
		}

		result := !playOptimal(n-1) || !playOptimal(n-2) || !playOptimal(n-3)
		memo[n] = result

		return result
	}

	return playOptimal(n)
}

// Time -> 0(1) Space -> 0(1)
func canWinNimI(n int) bool {
	return n%4 != 0
}

func main() {
	fmt.Println(canWinNim(16))
	fmt.Println(canWinNimI(16))
}
