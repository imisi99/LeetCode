package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head, tail := list1, list2
	if list1.Val > list2.Val {
		head, tail = list2, list1
	}

	merged := head
	mergedHead := merged
	head = head.Next

	for head != nil && tail != nil {
		if head.Val <= tail.Val {
			merged.Next = head
			head = head.Next
		} else {
			merged.Next = tail
			tail = tail.Next
		}
		merged = merged.Next
		merged.Next = nil
	}

	if head == nil && tail != nil {
		merged.Next = tail
	} else if tail == nil && head != nil {
		merged.Next = head
	}
	return mergedHead
}
