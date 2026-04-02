package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func reverseBetween(head *ListNode, left, right int) *ListNode {
	if left == right {
		return head
	}
	var prev *ListNode
	curr := head
	diff := left

	for left > 1 {
		prev = curr
		curr = curr.Next
		left--
	}

	var reversed *ListNode
	var endNode *ListNode
	for diff <= right {
		next := curr.Next
		curr.Next = reversed
		if reversed == nil {
			endNode = curr
		}
		reversed = curr
		curr = next
		diff++
	}

	if endNode != nil {
		endNode.Next = curr
	}

	if prev != nil {
		prev.Next = reversed
	} else {
		return reversed
	}

	return head
}
