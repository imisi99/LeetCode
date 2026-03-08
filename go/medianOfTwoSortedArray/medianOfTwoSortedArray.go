package main

import (
	"fmt"
	"math"
)

// Time -> 0(log(min(M, N))) Sapce -> 0(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	mid := (m + n + 1) / 2
	x, y := 0, m

	for x <= y {
		i := (x + y) / 2
		j := mid - i

		num1Left, num1Right := math.MinInt, math.MaxInt
		num2Left, num2Right := math.MinInt, math.MaxInt

		if i > 0 {
			num1Left = nums1[i-1]
		}
		if i < m {
			num1Right = nums1[i]
		}
		if j > 0 {
			num2Left = nums2[j-1]
		}
		if j < n {
			num2Right = nums2[j]
		}

		if num1Left <= num2Right && num2Left <= num1Right {
			maxLeft := max(num1Left, num2Left)
			if (m+n)%2 == 0 {
				minRight := min(num1Right, num2Right)
				return (float64(maxLeft) + float64(minRight)) / 2
			}
			return float64(maxLeft)
		}

		if num1Left > num2Right {
			y = i - 1
		} else {
			x = i + 1
		}
	}

	return 0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 12}, []int{4, 5, 6, 7, 8, 9}))
}
