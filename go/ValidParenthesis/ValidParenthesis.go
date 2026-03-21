package main

import "fmt"

func isvalid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		val := s[i]
		switch val {
		case '[':
			stack = append(stack, '[')
		case '(':
			stack = append(stack, '(')
		case '{':
			stack = append(stack, '{')
		default:
			var newVal byte
			if len(stack) > 0 {
				newVal = stack[len(stack)-1]
			} else {
				return false
			}
			switch newVal {
			case '[':
				if val != ']' {
					return false
				}
				stack = stack[:len(stack)-1]
			case '(':
				if val != ')' {
					return false
				}
				stack = stack[:len(stack)-1]
			case '{':
				if val != '}' {
					return false
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isvalid("[({})]"))
	fmt.Println(isvalid("[({))]"))
}
