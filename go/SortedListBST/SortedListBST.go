package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(N)
func sortedListToBST(head *ListNode) *TreeNode {
	var array []int
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	return createBST(array, 0, len(array)-1)
}

func createBST(array []int, start, stop int) *TreeNode {
	if start > stop {
		return nil
	}
	mid := (start + stop) / 2

	node := &TreeNode{Val: array[mid]}
	node.Left = createBST(array, start, mid-1)
	node.Right = createBST(array, mid+1, stop)

	return node
}
