package main

import "fmt"

func numDistinct(s, t string) int {
	memo := make(map[[2]int]int, 0)
	recurse(s, t, 0, 0, memo)
	return memo[[2]int{0, 0}]
}

func recurse(s, t string, si, ti int, memo map[[2]int]int) int {
	if si >= len(s) || ti >= len(t) {
		return 0
	}

	if val, exists := memo[[2]int{si, ti}]; exists {
		return val
	}

	count := 0
	if ti == len(t)-1 && s[si] == t[ti] {
		count++
	}

	if s[si] == t[ti] {
		count += recurse(s, t, si+1, ti+1, memo)
	}
	count += recurse(s, t, si+1, ti, memo)
	memo[[2]int{si, ti}] = count
	return count
}

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("babgbag", "bag"))
}
