package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(max(M, N)) Space -> 0(1)
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	carry := 0
	head := l1
	prev := l1
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		val := sum % 10
		carry = sum / 10
		l1.Val = val
		prev = l1
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil || l2 != nil {
		if l1 == nil {
			prev.Next = l2
			l1 = prev.Next
		}
		for carry != 0 && l1 != nil {
			sum := l1.Val + carry
			val := sum % 10
			carry = sum / 10
			l1.Val = val
			prev = l1
			l1 = l1.Next
		}
	}
	if carry != 0 {
		prev.Next = &ListNode{Val: carry}
	}

	return head
}
