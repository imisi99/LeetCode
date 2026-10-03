package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func deleteMiddle(head *ListNode) *ListNode {
	count := 0
	headI := head
	headII := head

	for head != nil {
		count++
		head = head.Next
	}

	mid := count / 2
	var prev *ListNode
	for mid != 0 {
		prev = headI
		headI = headI.Next
		mid--
	}

	if prev != nil {
		prev.Next = headI.Next
	}

	return headII
}
