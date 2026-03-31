package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(N)
func preorderTraversal(root *TreeNode) []int {
	array := []int{}
	preOrder(root, &array)
	return array
}

func preOrder(root *TreeNode, array *[]int) {
	if root == nil {
		return
	}

	*array = append(*array, root.Val)
	preOrder(root.Left, array)
	preOrder(root.Right, array)
}

func preorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	array := []int{}

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		array = append(array, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}

		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return array
}

func main() {
	node2 := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 2}
	root := &TreeNode{Val: 1, Left: node1, Right: node2}
	fmt.Println(preorderTraversal(root))
	fmt.Println(preorderTraversalIterative(root))
}
