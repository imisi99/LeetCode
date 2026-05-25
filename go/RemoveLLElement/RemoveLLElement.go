package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func removeElement(head *ListNode, val int) *ListNode {
	curr := head
	var prev *ListNode
	var start *ListNode
	for curr != nil {
		if curr.Val == val {
			if prev != nil {
				prev.Next = curr.Next
			}
			curr = curr.Next
		} else {
			if start == nil {
				start = curr
			}
			prev = curr
			curr = curr.Next
		}
	}
	return start
}
