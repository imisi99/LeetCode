package main

import (
	"fmt"
	"sort"
)

// Time => 0((k1+k2)*N) Space => 0(1)
func minSumSquareDiff(nums1, nums2 []int, k1, k2 int) int64 {
	for k1 != 0 || k2 != 0 {
		maxIdx := 0
		maxDiff := 0
		for idx := range nums1 {
			diff := (nums1[idx] - nums2[idx])
			diff *= diff

			if diff > maxDiff {
				maxDiff = diff
				maxIdx = idx
			}
		}

		if maxDiff == 0 {
			return 0
		}

		if k1 != 0 {
			if nums1[maxIdx] > nums2[maxIdx] {
				nums1[maxIdx]--
			} else {
				nums1[maxIdx]++
			}
			k1--
		} else if k2 != 0 {
			if nums2[maxIdx] > nums1[maxIdx] {
				nums2[maxIdx]--
			} else {
				nums2[maxIdx]++
			}
			k2--
		}
	}

	var sumDiff int64
	for idx := range nums1 {
		diff := (nums1[idx] - nums2[idx])
		diff *= diff
		sumDiff += int64(diff)
	}

	return sumDiff
}

// Time => 0(NlgoN) Space => 0(N)
func minSumSquareDiffI(nums1, nums2 []int, k1, k2 int) int64 {
	d := make([]int, len(nums2)+1)
	total := 0
	for idx := range nums1 {
		diff := (nums1[idx] - nums2[idx])
		if diff < 0 {
			diff = -diff
		}
		d[idx] = diff
		total += (diff * diff)
	}
	k := k1 + k2

	if total <= k {
		return 0
	}

	sort.Slice(d, func(i, j int) bool {
		return d[i] > d[j]
	})

	for i := range nums1 {
		group := (i + 1)
		cost := group * (d[i] - d[i+1])

		if k >= cost {
			k -= cost
			continue
		} else {
			q, r := k/group, k%group
			v := d[i] - q

			diff := ((group - r) * v * v) + ((v - 1) * (v - 1) * r)
			for j := i + 1; j < len(d); j++ {
				diff += (d[j] * d[j])
			}
			return int64(diff)
		}
	}

	return 0
}

func main() {
	fmt.Println(minSumSquareDiff([]int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0))
	fmt.Println(minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1))
	fmt.Println(minSumSquareDiffI([]int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0))
	fmt.Println(minSumSquareDiffI([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1))
}
