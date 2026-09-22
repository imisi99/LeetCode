package main

func resultArray(nums []int, k int, queries [][]int) []int {
	result := make([]int, len(queries))

	for qidx, query := range queries {
		idx, val, start, x := query[0], query[1], query[2], query[3]
		nums[idx] = val
		start = start - 1
		end := len(nums) - 1
		currSum := sum(start, end, nums)
		for end >= start {
			currSum /= nums[end]
			if currSum%k == x {
				result[qidx]++
			}
		}

	}

	return result
}

func sum(i, j int, array []int) int {
	prod := 1
	for i <= j {
		prod *= array[i]
		i++
	}

	return prod
}

func main() {}
