package main

import (
	"fmt"
	"strings"
)

// Time -> 0(N*M) Space -> 0(N*M)
func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	array := make([][]string, 0)

	for range numRows {
		subArray := []string{}
		for range s {
			subArray = append(subArray, "")
		}
		array = append(array, subArray)
	}

	sid := 0
	col := 0
	for sid < len(s) {
		for idx := range numRows {
			if sid == len(s) {
				break
			}
			array[idx][col] = string(s[sid])
			sid++
		}

		col++

		startRow := numRows - 2
		for startRow >= 1 && sid < len(s) {
			array[startRow][col] = string(s[sid])
			startRow--
			col++
			sid++
		}
	}

	newString := &strings.Builder{}
	row := 0
	for row < len(array) {
		col = 0
		for col < len(array[0]) {
			if array[row][col] != "" {
				newString.WriteString(array[row][col])
			}
			col++
		}
		row++
	}

	return newString.String()
}

// Time -> 0(N) Space -> 0(N)
func convertI(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	array := make([]strings.Builder, numRows)
	row, direction := 0, -1

	for _, ch := range s {
		array[row].WriteRune(ch)

		if row == 0 || row == numRows-1 {
			direction *= -1
		}

		row += direction
	}

	result := strings.Builder{}
	for _, builder := range array {
		result.WriteString(builder.String())
	}

	return result.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convertI("PAYPALISHIRING", 4))
}
