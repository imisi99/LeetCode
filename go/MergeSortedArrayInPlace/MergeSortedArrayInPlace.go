package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	end := m + n - 1
	m -= 1
	n -= 1

	for end >= 0 {
		if n < 0 {
			break
		}
		if m < 0 {
			for end >= 0 {
				nums1[end] = nums2[n]
				end--
				n--
			}
			break
		}
		if nums1[m] >= nums2[n] {
			nums1[end] = nums1[m]
			end--
			m--
		} else {
			nums1[end] = nums2[n]
			end--
			n--
		}
	}
}

func main() {
	array1 := []int{2, 3, 9, 0, 0, 0}
	array2 := []int{0, 1, 9}
	merge(array1, 3, array2, 3)
	fmt.Println(array1)
}
