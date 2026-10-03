package main

import "fmt"

// Time -> 0(N^2) Space -> 0(1)
func totalFruit(fruits []int) int {
	sum := 0

	for idx := range fruits {
		val := idx
		basket1 := -1
		basket2 := -1
		currSum := 0
		for val < len(fruits) {
			if basket1 == -1 {
				basket1 = fruits[val]
				currSum++
			} else if basket2 == -1 {
				basket2 = fruits[val]
				currSum++
			} else if fruits[val] == basket1 || fruits[val] == basket2 {
				currSum++
			} else {
				break
			}
			val++
		}

		if currSum > sum {
			sum = currSum
		}
	}
	return sum
}

func main() {
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
}
