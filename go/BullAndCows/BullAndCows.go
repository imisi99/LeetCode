package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func getHint(secret, guess string) string {
	lookup := make(map[byte]int, 0)

	for idx := range secret {
		lookup[secret[idx]]++
	}

	bull, cow := 0, 0

	for idx := range guess {
		val := guess[idx]
		if val == secret[idx] {
			lookup[val]--
			bull++
		}
	}

	for idx := range guess {
		if guess[idx] == secret[idx] {
			continue
		}
		val := guess[idx]
		if lookup[val] >= 1 {
			lookup[val]--
			cow++
		}
	}

	return fmt.Sprint(bull, "A", cow, "B")
}

func main() {
	fmt.Println(getHint("1807", "7810"))
	fmt.Println(getHint("1123", "0111"))
	fmt.Println(getHint("567", "567"))
	fmt.Println(getHint("11", "10"))
}
