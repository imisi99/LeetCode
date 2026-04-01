package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(M + N) Space -> 0(max(M, N))
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodes := make(map[*ListNode]bool, 0)

	for headA != nil {
		nodes[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if _, exists := nodes[headB]; exists {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func getIntersectionNodeI(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
