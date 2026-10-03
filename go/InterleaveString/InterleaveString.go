package main

import "fmt"

func interleave(a, b, s string) bool {
	left, right := make([]int, len(s)), make([]int, len(s))
	for i := range left {
		left[i], right[i] = -1, -1
	}

	recurse(a, b, s, 0, 0, 0, 0, left, right)
	recurse(a, b, s, 0, 0, 0, 1, left, right)

	fmt.Println(left, right)

	return left[0] == 1 || right[0] == 1
}

func recurse(a, b, s string, aPos, bPos, sPos, dir int, left, right []int) int {
	if sPos >= len(s) {
		if bPos >= len(b)-1 && aPos >= len(a)-1 {
			return 1
		} else {
			return 0
		}
	}

	if aPos == len(a) || bPos == len(b) {
		return 0
	}

	if dir == 0 {
		if a[aPos] != s[sPos] {
			return 0
		}
	} else {
		if b[bPos] != s[sPos] {
			return 0
		}
	}

	if left[sPos] == -1 {
		left[sPos] = recurse(a, b, s, aPos+1, bPos, sPos+1, 0, left, right)
	}
	if right[sPos] == -1 {
		right[sPos] = recurse(a, b, s, aPos, bPos+1, sPos+1, 1, left, right)
	}

	if left[sPos] == 1 {
		return 1
	}
	return right[sPos]
}

func main() {
	fmt.Println(interleave("a", "b", "ab"))
}
