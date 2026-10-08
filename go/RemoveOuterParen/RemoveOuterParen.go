package main

import (
	"fmt"
	"strings"
)

func removeOuterParentheses(s string) string {
	primitives := make([]string, 0)
	count := 0
	start := 0

	for idx := range s {
		if s[idx] == '(' {
			count++
		} else {
			count--
		}

		if count == 0 {
			primitives = append(primitives, s[start+1:idx])
			start = idx + 1
		}
	}

	return strings.Join(primitives, "")
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}
