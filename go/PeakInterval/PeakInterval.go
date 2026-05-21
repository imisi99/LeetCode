package main

import "math"

// Time -> 0(N) Space -> O(1)
func peak(intervals []int) int {
	start := math.MinInt

	for idx := range intervals {
		currVal := intervals[idx]
		end := 0
		if idx+1 == len(intervals) {
			end = math.MinInt
		} else {
			end = intervals[idx+1]
		}

		if currVal > start && currVal > end {
			return idx
		}

		start = currVal
	}
	return -1
}
