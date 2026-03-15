package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func strStr(haystack, needle string) int {
	idx := 0

	for idx < len(haystack) {
		if haystack[idx] == needle[0] {
			if idx+len(needle) <= len(haystack) {
				if haystack[idx:idx+len(needle)] == needle {
					return idx
				}
			} else {
				break
			}
		}
		idx++
	}
	return -1
}

func main() {
	fmt.Println(strStr("leetcode", "leetco"))
}
