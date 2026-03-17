package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(log(N))
func inorderTraversal(root *TreeNode) []int {
	array := []int{}
	inorder(root, &array)
	return array
}

func inorder(root *TreeNode, array *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, array)
	*array = append(*array, root.Val)
	inorder(root.Right, array)
}

func main() {
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 1, Left: left, Right: right}

	fmt.Println(inorderTraversal(root))
}
