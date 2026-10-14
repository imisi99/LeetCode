package main

import (
	"fmt"
	"strings"
)

// Time => 0(N) Space => 0(N)
func maxSumOfSquares(num, sum int) string {
	nines := sum / 9
	noNine := sum % 9

	count := nines
	if noNine > 0 {
		count++
	}
	if count > num {
		return ""
	}

	val := make([]string, num)

	i := 0
	for range nines {
		val[i] = "9"
		i++
	}
	if i < len(val) {
		val[i] = fmt.Sprint(noNine)
		i++

		for ; i < len(val); i++ {
			val[i] = "0"
		}
	}

	return strings.Join(val, "")
}

func main() {
	fmt.Println(maxSumOfSquares(3, 27))
}
