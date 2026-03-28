package main

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	return recurse(n)
}

// Time -> 0(N*M) Space -> 0(N + M)
func recurse(n int) string {
	if n == 1 {
		return "1"
	}

	rle := recurse(n - 1)

	builder := &strings.Builder{}

	idx := 0
	for idx < len(rle) {
		currIdx := idx
		for idx < len(rle) && rle[currIdx] == rle[idx] {
			idx++
		}
		length := fmt.Sprintf("%d", (idx - currIdx))
		builder.WriteString(length)
		builder.WriteByte(rle[currIdx])
	}

	return builder.String()
}

// Time -> 0(N*M) Space -> 0(M)
func approach1(n int) string {
	rle := "1"
	for i := 1; i < n; i++ {
		idx := 0
		builder := &strings.Builder{}
		for idx < len(rle) {
			currIdx := idx
			for idx < len(rle) && rle[currIdx] == rle[idx] {
				idx++
			}
			length := fmt.Sprintf("%d", (idx - currIdx))
			builder.WriteString(length)
			builder.WriteByte(rle[currIdx])
		}

		rle = builder.String()
	}

	return rle
}

func main() {
	fmt.Println(countAndSay(10))
	fmt.Println(approach1(10))
}
