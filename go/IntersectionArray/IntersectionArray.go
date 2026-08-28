package main

import "fmt"

// Time -> 0(N + M), Space -> 0(min(M, N))
func intersection(nums1, nums2 []int) []int {
	lookup := make(map[int]bool, 0)

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for idx := range nums1 {
		lookup[nums1[idx]] = true
	}

	result := make([]int, 0)

	for idx := range nums2 {
		if _, exists := lookup[nums2[idx]]; exists {
			result = append(result, nums2[idx])
			delete(lookup, nums2[idx])
		}
	}

	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 3, 2, 4, 9}, []int{9, 4, 2}))
}
