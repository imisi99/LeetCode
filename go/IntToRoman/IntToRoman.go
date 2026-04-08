package main

import (
	"fmt"
	"strings"
)

// Time -> 0(N)
func intToRoman(num int) string {
	result := &strings.Builder{}
	for num > 0 {
		if num/1000 >= 1 {
			result.WriteString("M")
			num -= 1000
		} else if num/100 >= 1 {
			if num >= 400 && num <= 499 {
				result.WriteString("CD")
				num -= 400
			} else if num >= 900 && num <= 999 {
				result.WriteString("CM")
				num -= 900
			} else if num >= 500 {
				result.WriteString("D")
				num -= 500
			} else {
				result.WriteString("C")
				num -= 100
			}
		} else if num/10 >= 1 {
			if num >= 40 && num <= 49 {
				result.WriteString("XL")
				num -= 40
			} else if num >= 90 && num <= 99 {
				result.WriteString("XC")
				num -= 90
			} else if num >= 50 {
				result.WriteString("L")
				num -= 50
			} else {
				result.WriteString("X")
				num -= 10
			}
		} else {
			if num == 4 {
				result.WriteString("IV")
				num -= 4
			} else if num == 9 {
				result.WriteString("IX")
				num -= 9
			} else if num >= 5 {
				result.WriteString("V")
				num -= 5
			} else {
				result.WriteString("I")
				num -= 1
			}
		}
	}
	return result.String()
}

func main() {
	fmt.Println(intToRoman(3979))
}
