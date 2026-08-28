package main

import "fmt"

// Time -> 0(N + M), Space -> 0(min(M, N))
func intersect(nums1, nums2 []int) []int {
	lookup := make(map[int]int, 0)

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for idx := range nums1 {
		lookup[nums1[idx]]++
	}

	result := make([]int, 0)

	for idx := range nums2 {
		if count, exists := lookup[nums2[idx]]; exists && count > 0 {
			result = append(result, nums2[idx])
			lookup[nums2[idx]]--
		}
	}

	return result
}

func main() {
	fmt.Println(intersect([]int{1, 2, 3, 2, 4, 9}, []int{2, 9, 4, 2, 2}))
}
