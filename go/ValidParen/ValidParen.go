package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func isValid(s string) bool {
	stack := make([]byte, 0)

	for idx := range s {
		switch s[idx] {
		case '(':
			stack = append(stack, '(')
		case '[':
			stack = append(stack, '[')
		case '{':
			stack = append(stack, '{')
		case ')':
			val := popStack(stack)
			if val != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			val := popStack(stack)
			if val != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			val := popStack(stack)
			if val != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func popStack(stack []byte) byte {
	if len(stack) == 0 {
		return '8'
	}
	return stack[len(stack)-1]
}

func main() {
	fmt.Println(isValid("()"))
}
