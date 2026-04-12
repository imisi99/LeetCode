package main

import "fmt"

// Time -> 0(N*N) Space -> 0(N*N)
func generate(numRows int) [][]int {
	result := make([][]int, 0)

	for i := 1; i <= numRows; i++ {
		switch i {
		case 1:
			result = append(result, []int{1})
		case 2:
			result = append(result, []int{1, 1})
		default:
			prevRow := result[i-2]
			newRow := make([]int, i)
			newRow[0], newRow[i-1] = 1, 1

			idx := 1
			for i := 0; i < len(prevRow)-1; i++ {
				newRow[idx] = prevRow[i] + prevRow[i+1]
				idx++
			}
			result = append(result, newRow)
		}
	}
	return result
}

func main() {
	fmt.Println(generate(5))
}
