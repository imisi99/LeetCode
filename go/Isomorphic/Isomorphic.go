package main

import "fmt"

// Time -> 0(N) Space -> 0(N)
func isIsomorphic(s, t string) bool {
	mapped := make(map[byte]byte, 0)

	if len(s) != len(t) {
		return false
	}

	for i := range s {
		if val, exists := mapped[s[i]]; exists {
			if val != t[i] {
				return false
			}
		}
		mapped[s[i]] = t[i]
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("add", "egg"))
	fmt.Println(isIsomorphic("f11", "egd"))
}
