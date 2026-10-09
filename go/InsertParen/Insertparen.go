package main

import "fmt"

// Time => 0(N) Space => 0(1)
func minInsertions(s string) int {
	stack := 0

	waiting := false
	insertions := 0
	for idx := 0; idx < len(s); idx++ {
		if s[idx] == '(' {
			if waiting {
				if stack == 0 {
					insertions++
				} else {
					stack--
				}
				insertions++
			}
			stack++
			waiting = false
		} else {
			if waiting {
				if stack == 0 {
					insertions++
				} else {
					stack--
				}
				waiting = false
			} else {
				waiting = true
			}
		}
	}

	if stack > 0 {
		if waiting {
			insertions++
			stack--
		}
		insertions += (stack * 2)
	} else {
		if waiting {
			insertions += 2
		}
	}

	return insertions
}

func main() {
	fmt.Println(minInsertions("(()))"))
	fmt.Println(minInsertions("())"))
	fmt.Println(minInsertions("))())("))
}
