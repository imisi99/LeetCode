package main

import "fmt"

// Time -> 0(N) Space -> 0(N) where N is the no of reserved seats
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	array := make([][]int, n)

	for _, pos := range reservedSeats {
		array[pos[0]-1] = append(array[pos[0]-1], pos[1])
	}

	count := 0
	for _, pos := range array {
		block1 := true
		block2 := true
		block3 := true

		for _, idx := range pos {
			if idx >= 2 && idx <= 5 {
				block1 = false
			}
			if idx >= 4 && idx <= 7 {
				block2 = false
			}
			if idx >= 6 && idx <= 9 {
				block3 = false
			}
		}

		if block1 || block2 || block3 {
			if block1 && block3 {
				count += 2
			} else {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(maxNumberOfFamilies(4, [][]int{
		{4, 3}, {1, 4}, {4, 6}, {1, 7},
	}))
}
