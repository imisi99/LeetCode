package main

import (
	"fmt"
)

// Time -> 0(N^3)
func longestPalindrome(s string) string {
	longest := []int{0, 0}

	for idx := range s {
		lasts := []int{}
		for j := len(s) - 1; j > idx; j-- {
			if s[idx] == s[j] {
				lasts = append(lasts, j)
			}
		}

		for _, last := range lasts {
			start, stop := idx, last
			for start < stop && s[start] == s[stop] {
				start++
				stop--
			}

			if start == stop || s[start] == s[stop] {
				if (last - idx) > (longest[1] - longest[0]) {
					longest[1], longest[0] = last, idx
				}
			}

		}
	}

	return s[longest[0] : longest[1]+1]
}

// Time -> 0(N^2)
func expandCenterPalindrome(s string) string {
	expand := func(i, j int) (int, int) {
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i--
			j++
		}
		return i + 1, j - 1
	}

	start, stop := 0, 0
	for i := range s {
		l1, r1 := expand(i, i)
		l2, r2 := expand(i, i+1)

		if r1-l1 > stop-start {
			stop, start = r1, l1
		}

		if r2-l2 > stop-start {
			stop, start = r2, l2
		}
	}

	return s[start : stop+1]
}

func main() {
	fmt.Println(longestPalindrome("Gbbaabd"))
	fmt.Println(expandCenterPalindrome("Gbbaabd"))
}
