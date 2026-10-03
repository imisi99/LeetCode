package main

import "fmt"

// Time -> 0(N*M) Space -> 0(M)
func resultArray(nums []int, k int, queries [][]int) []int {
	result := make([]int, len(queries))

	for qidx, query := range queries {
		idx, val, start, x := query[0], query[1], query[2], query[3]
		nums[idx] = val
		prod := 1
		for end := start; end < len(nums); end++ {
			prod = (prod * (nums[end] % k)) % k

			if prod == x {
				result[qidx]++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(resultArray([]int{1, 2, 3, 4, 5}, 3, [][]int{{2, 2, 0, 2}, {3, 3, 3, 0}, {0, 1, 0, 1}}))
}
