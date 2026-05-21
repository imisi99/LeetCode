package main

import (
	"fmt"
	"strings"
)

// Time -> 0(N) Space -> (N)
func wordPattern(pattern, s string) bool {
	patMap := make(map[byte]int, 0)
	sMap := make(map[string]int, 0)

	words := strings.Split(s, " ")

	for pat := range pattern {
		patMap[pattern[pat]]++
	}

	for word := range words {
		sMap[words[word]]++
	}

	pMap := make(map[int]int, 0)
	wMap := make(map[int]int, 0)

	for _, val := range patMap {
		pMap[val]++
	}

	for _, val := range sMap {
		wMap[val]++
	}

	if len(pMap) != len(wMap) {
		return false
	}

	for key, val := range pMap {
		if wVal, exists := wMap[key]; !exists || wVal != val {
			return false
		}
	}
	return true
}

// Time -> 0(N) Space -> 0(N)
func wordPatternII(pattern, s string) bool {
	letterMap := make(map[byte]string, 0)
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	for pat := range pattern {
		if val, exists := letterMap[pattern[pat]]; exists && val != words[pat] {
			return false
		}
		letterMap[pattern[pat]] = words[pat]
	}
	return true
}

func main() {
	fmt.Println(wordPattern("aba", "dog cat cat"))
	fmt.Println(wordPatternII("aba", "dog cat cat"))
}
