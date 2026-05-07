package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func partition(head *ListNode, x int) *ListNode {
	curr := head
	var lessHead, less, prev, pivot *ListNode
	for curr != nil {
		if curr.Val >= x {
			pivot = curr
			currPrev := prev
			for curr != nil {
				if curr.Val < x {
					if lessHead == nil {
						less = curr
						lessHead = less
					} else {
						less.Next = curr
						less = less.Next
					}
					currPrev.Next = curr.Next
				} else {
					currPrev = curr
				}
				curr = curr.Next
			}
			break
		}
		prev = curr
		curr = curr.Next
	}

	if less != nil {
		if prev == nil {
			less.Next = head
			return lessHead
		}

		prev.Next = lessHead
		less.Next = pivot
	}
	return head
}

func main() {
	head := &ListNode{}
	currHead := head
	for _, val := range []int{2, 3, 4, 3, 1, 2, 5} {
		currHead.Val = val
		currHead.Next = &ListNode{}
		currHead = currHead.Next
	}

	res := partition(head, 3)
	var array []int
	for res != nil {
		array = append(array, res.Val)
		res = res.Next
	}
	fmt.Println(array)
}
