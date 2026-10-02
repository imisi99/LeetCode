package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	array := make([]string, 0)
	recurse(0, n, []byte{}, &array)
	return array
}

func recurse(open, n int, trace []byte, array *[]string) {
	if open < 0 || open > n {
		return
	}

	if len(trace) == (2 * n) {
		if open == 0 {
			*array = append(*array, string(trace))
		}
		return
	}

	trace = append(trace, '(')
	recurse(open+1, n, trace, array)

	trace[len(trace)-1] = ')'
	recurse(open-1, n, trace, array)
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
