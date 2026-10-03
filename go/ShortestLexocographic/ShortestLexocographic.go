package main

import "fmt"

func smallest(s string, k int) string {
	array := make([]string, 0)
	lookup := make(map[int][]int, 0)

	small := len(s)
	idx := 0
	for i := range s {
		j := i
		ones := 0
		for j < len(s) {
			if string(s[j]) == "1" {
				ones++
				if ones == k {
					newLen := j - i + 1
					array = append(array, string(s[i:j+1]))
					lookup[newLen] = append(lookup[newLen], idx)
					idx++
					if newLen < small {
						small = newLen
					}
				}
			}

			j++
		}
	}

	if len(array) == 0 {
		return ""
	}

	zeros := lookup[small]
	if len(zeros) == 1 {
		return array[zeros[0]]
	}

	idx = 0
	for idx < small {
		newZero := make([]int, 0)
		ones := make([]int, 0)
		for _, id := range zeros {
			if string(array[id][idx]) == "0" {
				newZero = append(newZero, id)
			} else {
				ones = append(ones, id)
			}
		}
		if len(newZero) == 1 {
			return array[newZero[0]]
		} else if len(newZero) == 0 {
			newZero = ones
		}
		idx++

		zeros = newZero
	}

	return array[zeros[0]]
}

func main() {
	fmt.Println(smallest("100011001", 3))
}
