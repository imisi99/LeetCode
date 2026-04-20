package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flattenBinaryTree(root *TreeNode) {
	if root == nil {
		return
	}
	convLinkedList(root)
}

// Time -> 0(N) Space -> 0(log(N))
func convLinkedList(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right == nil {
		return root
	}

	var storedRight *TreeNode
	if root.Left != nil {
		storedRight = root.Right
		root.Right, root.Left = root.Left, nil
		root = root.Right
		root = convLinkedList(root)
		root.Right = storedRight
	}

	if root.Right != nil {
		root = root.Right
		root = convLinkedList(root)
	}
	return root
}

func printList(root *TreeNode) {
	array := make([]int, 0)
	for root != nil {
		array = append(array, root.Val)
		root = root.Right
	}
	fmt.Println(array)
}

func main() {
	a := &TreeNode{Val: 4}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3, Right: a}
	root := &TreeNode{Val: 1, Left: left, Right: right}

	flattenBinaryTree(root)
	printList(root)
}
