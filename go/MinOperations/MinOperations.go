package main

import "fmt"

func minOperations(nums []int, x int) int {
	top := recurse(x, 0, len(nums)-1, 0, true, nums)
	bot := recurse(x, 0, len(nums)-1, 0, false, nums)

	return optimal(top, bot)
}

func recurse(x, topIdx, botIdx, moves int, top bool, nums []int) int {
	if x <= 0 {
		if x < 0 {
			return -1
		}
		return moves
	}

	if botIdx < topIdx || botIdx < 0 || topIdx == len(nums) {
		return -1
	}

	if top {
		x -= nums[topIdx]
	} else {
		x -= nums[botIdx]
	}

	topMoves := recurse(x, topIdx+1, botIdx, moves+1, true, nums)
	botMoves := recurse(x, topIdx, botIdx-1, moves+1, false, nums)

	return optimal(topMoves, botMoves)
}

func optimal(i, j int) int {
	if i == -1 {
		return j
	} else if j == -1 {
		return i
	} else {
		return min(i, j)
	}
}

func main() {
	fmt.Println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	fmt.Println(minOperations([]int{5, 6, 7, 8, 9}, 4))
	fmt.Println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
}
