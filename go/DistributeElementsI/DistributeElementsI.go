package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func resultArray(nums []int) []int {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	arr1 = append(arr1, nums[0])
	arr2 = append(arr2, nums[1])

	idx := 2
	for idx < len(nums) {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[idx])
		} else {
			arr2 = append(arr2, nums[idx])
		}
		idx++
	}

	arr1 = append(arr1, arr2...)

	return arr1
}

func main() {
	fmt.Println(resultArray([]int{2, 1, 3}))
}
