package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0

	builder := &strings.Builder{}
	for i >= 0 && j >= 0 {
		val := carry + int(a[i]-'0') + int(b[j]-'0')
		switch val {
		case 3:
			builder.WriteString("1")
			carry = 1
		case 2:
			builder.WriteString("0")
			carry = 1
		case 1:
			builder.WriteString("1")
			carry = 0
		default:
			builder.WriteString("0")
			carry = 0
		}
		i--
		j--
	}

	for i >= 0 {
		val := carry + int(a[i]-'0')
		if val == 2 {
			builder.WriteString("0")
			carry = 1
		} else {
			builder.WriteString(strconv.Itoa(val))
			carry = 0
		}
		i--
	}

	for j >= 0 {
		val := carry + int(b[j]-'0')
		if val == 2 {
			builder.WriteString("0")
			carry = 1
		} else {
			builder.WriteString(strconv.Itoa(val))
			carry = 0
		}
		j--
	}

	if carry == 1 {
		builder.WriteString("1")
	}

	return reverse(builder.String())
}

func reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
