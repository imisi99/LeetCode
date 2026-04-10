package main

import "fmt"

// Time -> 0(N) Space -> 0(1)
func lenghtOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	set := make(map[rune][2]int)

	startIdx := 0
	bestLen := 0
	for idx, ch := range s {
		val := set[ch]
		if val[0] > 0 && val[1] >= startIdx {
			newLen := idx - startIdx
			if newLen > bestLen {
				bestLen = newLen
			}
			startIdx = val[1] + 1
		}
		set[ch] = [2]int{1, idx}
	}

	if len(s)-startIdx > bestLen {
		return len(s) - startIdx
	}
	return bestLen
}

func main() {
	fmt.Println(lenghtOfLongestSubstring("bbbbb"))
}
