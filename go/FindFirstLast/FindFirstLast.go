package main

import (
	"fmt"
)

// Time -> 0(N) Space -> 0(1)
func linearSearchRange(nums []int, target int) []int {
	result := make([]int, 2)
	result[0], result[1] = -1, -1

	for i := range nums {
		if nums[i] == target {
			if result[0] == -1 {
				result[0] = i
			}
			result[1] = i
		}
	}

	return result
}

// Time -> 0(log N) Space -> 0(1)
func binarySearchRange(nums []int, target int) []int {
	result := make([]int, 2)
	result[0], result[1] = -1, -1

	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			tmpend := mid - 1
			tmpstart := mid + 1
			result[0], result[1] = mid, mid

			for start <= tmpend {
				mid = (start + tmpend) / 2
				if nums[mid] < target {
					start = mid + 1
				} else {
					result[0] = mid
					tmpend = mid - 1
				}
			}

			for tmpstart <= end {
				mid = (tmpstart + end) / 2
				if nums[mid] > target {
					end = mid - 1
				} else {
					result[1] = mid
					tmpstart = mid + 1
				}
			}

			break
		}
	}

	return result
}

func main() {
	array := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(linearSearchRange(array, 8))
	fmt.Println(binarySearchRange(array, 8))
}
