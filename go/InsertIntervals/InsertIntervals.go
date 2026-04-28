package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func insertIntervals(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	insertedInterval := make([][]int, 0)

	for idx := range intervals {
		if intervals[idx][0] > newInterval[0] {
			insertedInterval = append(insertedInterval, newInterval)
			insertedInterval = append(insertedInterval, intervals[idx:]...)
			break
		}
		insertedInterval = append(insertedInterval, intervals[idx])
	}

	if len(intervals) == len(insertedInterval) {
		insertedInterval = append(insertedInterval, newInterval)
	}

	mergedInterval := make([][]int, 0)

	i := 0
	for i < len(insertedInterval) {
		start := i
		i++
		for i < len(insertedInterval) && insertedInterval[i][0] <= insertedInterval[start][1] {
			insertedInterval[start][1] = max(insertedInterval[i][1], insertedInterval[start][1])
			i++
		}
		mergedInterval = append(mergedInterval, insertedInterval[start])
	}

	return mergedInterval
}

func main() {
	fmt.Println(insertIntervals([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{4, 8}))
	fmt.Println(insertIntervals([][]int{
		{1, 5},
	}, []int{3, 7}))
}
