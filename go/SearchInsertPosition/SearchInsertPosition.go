package main

import "fmt"

// Time -> 0(log(N)) Space -> 0(1)
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}

		if start > end {
			return start
		}
	}
	return 0
}

func main() {
	array := []int{1, 3, 5, 7, 9}
	fmt.Println(searchInsert(array, 4))
}
