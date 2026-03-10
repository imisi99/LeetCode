package main

import "fmt"

type LinkedList struct {
	Val  int
	Next *LinkedList
}

// Time -> 0(N) Space -> 0(N)
func removeduplicates(head *LinkedList) *LinkedList {
	set := make(map[int]int, 0)
	var newHead *LinkedList
	actualHead := newHead
	tmpHead := head

	for head != nil {
		set[head.Val]++
		head = head.Next
	}

	head = tmpHead
	for head != nil {
		tmpVal := head.Val
		if val := set[head.Val]; val > 1 {
			for head != nil && head.Val == tmpVal {
				head = head.Next
			}
		} else {
			if newHead == nil {
				newHead = head
				actualHead = newHead
				head = head.Next
			} else {
				newHead.Next = head
				head = head.Next
				newHead = newHead.Next
			}
			newHead.Next = nil
		}
	}

	return actualHead
}

func main() {
	B := &LinkedList{Val: 2, Next: nil}
	A := &LinkedList{Val: 1, Next: B}
	head := &LinkedList{Val: 1, Next: A}
	fmt.Println(removeduplicates(head))
}
