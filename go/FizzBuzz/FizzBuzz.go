package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func fizzBuzz(n int) []string {
	result := make([]string, 0)
	idx := 1
	for idx <= n {
		fizz := idx % 3
		buzz := idx % 5
		switch {
		case buzz == fizz && buzz == 0:
			result = append(result, "FizzBuzz")
		case fizz == 0:
			result = append(result, "Fizz")
		case buzz == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, fmt.Sprintf("%d", idx))
		}

		idx++
	}

	return result
}

func main() {
	fmt.Println(fizzBuzz(3))
	fmt.Println(fizzBuzz(5))
	fmt.Println(fizzBuzz(15))
}
