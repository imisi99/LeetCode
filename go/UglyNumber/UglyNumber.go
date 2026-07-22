package main

import "fmt"

func ugly(num int) bool {
	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else if num%3 == 0 {
			num = num / 3
		} else if num%5 == 0 {
			num = num / 5
		} else {
			break
		}
	}
	return num == 1
}

func main() {
	fmt.Println(ugly(5))
	fmt.Println(ugly(7))
}
