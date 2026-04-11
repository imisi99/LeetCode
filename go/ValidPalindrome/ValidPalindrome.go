package main

import (
	"fmt"
)

// Time -> 0(N) Space -> 0(1)
func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		for start < end && !isAlphaNum(s[start]) {
			start++
		}

		for end > start && !isAlphaNum(s[end]) {
			end--
		}

		if ToLower(s[start]) != ToLower(s[end]) {
			return false
		}
		start++
		end--
	}

	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func ToLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
