package main

import (
	"fmt"
)

func RotateArray(array []int, k int) {
	approachII(array, k)
}

// Time -> 0(KN) Space -> 0(1)
func approachI(array []int, k int) {
	for range k {
		last := array[len(array)-1]
		start := len(array) - 2
		for start >= 0 {
			array[start], array[start+1] = array[start+1], array[start]
			start--
		}
		array[0] = last
	}
}

// Time -> 0(N) Space -> 0(1)
func approachII(array []int, k int) {
	n := len(array)
	k = k % len(array)
	cycles := gcd(n, k)
	perCycle := n / cycles

	for start := range cycles {
		next, val := start, array[start]
		for range perCycle {
			pos := (next + k) % n
			tmp := array[pos]
			array[pos] = val
			next = pos
			val = tmp
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	arrayI := []int{1, 2, 3, 4, 5, 6, 7}

	approachI(array, 10)
	approachII(arrayI, 10)
	fmt.Println(array)
	fmt.Println(arrayI)
}
