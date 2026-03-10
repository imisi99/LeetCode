package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func deleteDuplicates(head *ListNode) *ListNode {
	stillHead := head
	for head != nil {
		currHead := head
		head = head.Next
		for head != nil && currHead.Val == head.Val {
			head = head.Next
		}
		currHead.Next = head
	}
	return stillHead
}
