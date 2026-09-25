package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func findTheDifference(s, t string) byte {
	lookup := make(map[byte]int, 0)

	for idx := range s {
		lookup[s[idx]]++
	}

	for idx := range t {
		if val, exist := lookup[t[idx]]; exist && val > 0 {
			lookup[t[idx]]--
		} else {
			return t[idx]
		}
	}

	return 0
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
	fmt.Println(findTheDifference("", "y"))
}
