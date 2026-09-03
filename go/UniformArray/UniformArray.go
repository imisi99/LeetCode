package main

import "fmt"

func uniformArray(nums1 []int) bool {
	lookup := make(map[int]bool, 0)
	minVal := nums1[0]

	for idx := range nums1 {
		lookup[nums1[idx]] = true
	}

	// Odd
	for idx := range nums1 {
		if nums1[idx]%2 == 0 {
			val := nums1[idx] - 1
			if exists := lookup[val]; !exists {
				odd := 2
				found := false
				for !found && val-odd >= minVal {
					found = lookup[val-odd]
					odd += 2
				}
				if !found {
					break
				}
			}
		}
		if idx == len(nums1)-1 {
			return true
		}
	}

	// Even
	for idx := range nums1 {
		if nums1[idx]%2 == 1 {
			val := nums1[idx] - 1
			if exists := lookup[val]; !exists {
				even := 2
				found := false
				for !found && val-even >= minVal {
					found = lookup[val-even]
					even += 2
				}
				if !found {
					break
				}
			}
		}
		if idx == len(nums1)-1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(uniformArray([]int{1, 7, 3}))
	fmt.Println(uniformArray([]int{4, 6}))
	fmt.Println(uniformArray([]int{2, 3}))
}
