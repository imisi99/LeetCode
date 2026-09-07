package main

import (
	"fmt"
	"strings"
)

func distinctSubSeqII(s string) int {
	distinct := make(map[string]bool, 0)
	recurse(s, 0, &strings.Builder{}, distinct)
	return len(distinct)
}

func recurse(s string, idx int, builder *strings.Builder, distinct map[string]bool) {
	if builder.Len() > 0 {
		if exists := distinct[builder.String()]; !exists {
			distinct[builder.String()] = true
		}
	}

	if idx == len(s) {
		return
	}

	newBuf := &strings.Builder{}
	newBufI := &strings.Builder{}
	newBuf.WriteString(builder.String())
	newBuf.WriteByte(s[idx])
	newBufI.WriteString(builder.String())

	recurse(s, idx+1, newBuf, distinct)
	recurse(s, idx+1, newBufI, distinct)
}

func main() {
	fmt.Println(distinctSubSeqII("abc"))
	fmt.Println(distinctSubSeqII("aba"))
	fmt.Println(distinctSubSeqII("aaa"))
}
