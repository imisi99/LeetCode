package main

import (
	"fmt"
	"strings"
)

func letterCombinations(digits string) []string {
	var result []string

	set := make(map[byte][]string, 0)
	set['2'] = []string{"a", "b", "c"}
	set['3'] = []string{"d", "e", "f"}
	set['4'] = []string{"g", "h", "i"}
	set['5'] = []string{"j", "k", "l"}
	set['6'] = []string{"m", "n", "o"}
	set['7'] = []string{"p", "q", "r", "s"}
	set['8'] = []string{"t", "u", "v"}
	set['9'] = []string{"w", "x", "y", "z"}

	builder := &strings.Builder{}
	traverse(builder, 0, digits, set, &result)
	return result
}

func traverse(builder *strings.Builder, idx int, digits string, set map[byte][]string, result *[]string) {
	if idx == len(digits) {
		return
	}

	for _, digit := range set[digits[idx]] {
		newBuilder := &strings.Builder{}
		newBuilder.WriteString(builder.String())
		newBuilder.WriteString(digit)
		traverse(newBuilder, idx+1, digits, set, result)
		if idx == len(digits)-1 {
			*result = append(*result, newBuilder.String())
		}
	}
}

func main() {
	fmt.Println(letterCombinations("27"))
}
