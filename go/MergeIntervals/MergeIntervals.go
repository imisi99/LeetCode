package main

import (
	"fmt"
	"sort"
)

// Time -> 0(NlogN) Space -> 0(N)
func merge(intervals [][]int) [][]int {
	merged := make([][]int, 0)
	idx := 0

	fmt.Println(intervals)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	fmt.Println(intervals)

	for idx < len(intervals) {
		start := idx
		idx++
		for idx < len(intervals) && intervals[start][1] >= intervals[idx][0] {
			intervals[start][1] = max(intervals[start][1], intervals[idx][1])
			idx++
		}
		merged = append(merged, intervals[start])
	}
	return merged
}

func main() {
	fmt.Println(merge([][]int{
		{4, 7},
		{1, 4},
	}))
}
