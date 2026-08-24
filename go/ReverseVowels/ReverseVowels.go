package main

import (
	"fmt"
	"strings"
)

// Time -> 0(N) Space -> 0(N)
func reverseVowels(s string) string {
	vowels := make([]byte, 0)

	for idx := range s {
		if strings.ContainsAny(string(s[idx]), "aeiouAEIOU") {
			vowels = append(vowels, s[idx])
		}
	}

	reversed := &strings.Builder{}

	for idx := range s {
		if strings.ContainsAny(string(s[idx]), "aeiouAEIOU") {
			reversed.WriteByte(vowels[len(vowels)-1])
			vowels = vowels[:len(vowels)-1]
		} else {
			reversed.WriteByte(s[idx])
		}
	}

	return reversed.String()
}

func main() {
	fmt.Println(reverseVowels("IceCreAm"))
}
