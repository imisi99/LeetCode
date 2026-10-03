package main

import "fmt"

func minMoves(classroom []string, energy int) int {
	pos1, pos2, liter := 0, 0, 0
	for i := range classroom {
		for j := range classroom[i] {
			switch classroom[i][j] {
			case 'L':
				liter++
			case 'S':
				pos1, pos2 = i, j
			}
		}
	}
	fmt.Println(pos1, pos2, liter)
	return recurse(classroom, energy, energy, liter, pos1, pos2, 0)
}

func recurse(classroom []string, maxEnergy, energy, liter, pos1, pos2, moves int) int {
	if liter == 0 {
		return moves
	}

	if pos1 < 0 || pos1 >= len(classroom) || pos2 < 0 || pos2 >= len(classroom[pos1]) {
		return -1
	}

	if energy == 0 && classroom[pos1][pos2] != 'R' {
		return -1
	}

	moves++
	energy--

	switch classroom[pos1][pos2] {
	case 'R':
		energy = maxEnergy
	case 'L':
		liter--
	case 'X':
		return -1
	}

	var moveCount [4]int
	moveCount[0] = recurse(classroom, maxEnergy, energy, liter, pos1+1, pos2, moves)
	moveCount[1] = recurse(classroom, maxEnergy, energy, liter, pos1-1, pos2, moves)
	moveCount[2] = recurse(classroom, maxEnergy, energy, liter, pos1, pos2-1, moves)
	moveCount[3] = recurse(classroom, maxEnergy, energy, liter, pos1, pos2+1, moves)

	minMove := -1
	for _, count := range moveCount {
		if count != -1 {
			if minMove == -1 {
				minMove = count
			} else {
				minMove = min(minMove, count)
			}
		}
	}
	return minMove
}

func main() {
	fmt.Println(minMoves([]string{"S.", "XL"}, 2))
	fmt.Println(minMoves([]string{"LS", "RL"}, 4))
	fmt.Println(minMoves([]string{"L.S", "RXL"}, 3))
}
