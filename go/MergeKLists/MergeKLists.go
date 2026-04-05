package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N*Max(M)) Space -> 0(1)
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 1 {
		if len(lists) == 0 {
			return nil
		}
		return lists[0]
	}
	mergedHead := lists[0]
	for _, list := range lists[1:] {
		merged := mergedHead
		var mergedPrev *ListNode
		for list != nil && merged != nil {
			if list.Val < merged.Val {
				listNext := list.Next
				if mergedPrev != nil {
					mergedPrev.Next = list
				} else {
					mergedHead = list
				}
				mergedPrev = list
				list.Next = merged
				list = listNext
			} else {
				mergedPrev = merged
				merged = merged.Next
			}
		}

		if list != nil {
			if mergedPrev != nil {
				mergedPrev.Next = list
			} else {
				mergedHead = list
			}
		}

	}
	return mergedHead
}
