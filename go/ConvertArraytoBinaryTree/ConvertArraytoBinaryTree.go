package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return addBinaryNode(nums, 0, len(nums)-1)
}

// Time -> 0(N) Space -> 0(log(N))
func addBinaryNode(nums []int, start, stop int) *TreeNode {
	if start > stop {
		return nil
	}

	mid := (start + stop) / 2
	node := &TreeNode{Val: nums[mid]}

	node.Left = addBinaryNode(nums, start, mid-1)
	node.Right = addBinaryNode(nums, mid+1, stop)

	return node
}

func viewTreeNodeInOrder(root *TreeNode) {
	if root == nil {
		return
	}

	viewTreeNodeInOrder(root.Left)
	fmt.Println(root.Val)
	viewTreeNodeInOrder(root.Right)
}

func main() {
	root := sortedArrayToBST([]int{-4, -1, 0, 1, 4})
	viewTreeNodeInOrder(root)
}
