package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func firstUniqChar(s string) int {
	occur := make(map[byte]int, 0)

	for char := range s {
		occur[s[char]]++
	}

	for char := range s {
		if occur[s[char]] == 1 {
			return char
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("loveleetcode"))
}
