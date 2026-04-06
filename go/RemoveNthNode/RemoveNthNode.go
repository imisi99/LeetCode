package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head

	k := 0
	for current != nil {
		current = current.Next
		k++
	}

	if k == n {
		return head.Next
	}

	stop := k - n - 1

	current = head
	for stop != 0 {
		current = current.Next
		stop--
	}

	current.Next = current.Next.Next

	return head
}
