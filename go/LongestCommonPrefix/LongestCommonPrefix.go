// Package longestcommonprefix
package main

import (
	"fmt"
)

// longestcommonprefix Time -> 0(N*M) where M is the length of the longest word Space -> 0(1)
func longestcommonprefix(array []string) string {
	prefix := array[0]
	for _, val := range array[1:] {
		for i := range prefix {
			if i == len(val) || prefix[i] != val[i] {
				prefix = prefix[:i]
				break
			}
		}
	}
	return prefix
}

func main() {
	array := []string{"profile", "prolific", "pro", "property"}
	prefix := longestcommonprefix(array)
	fmt.Println(prefix)
}
