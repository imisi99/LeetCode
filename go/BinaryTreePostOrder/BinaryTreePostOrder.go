package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	array := []int{}
	postOrder(root, &array)
	return array
}

func postOrder(root *TreeNode, array *[]int) {
	if root == nil {
		return
	}

	postOrder(root.Left, array)
	postOrder(root.Right, array)
	*array = append(*array, root.Val)
}

func postorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	array := []int{}
	stack := []*TreeNode{}

	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		array = append(array, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}

	return array
}

func main() {
	node2 := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 2}
	root := &TreeNode{Val: 1, Left: node1, Right: node2}
	fmt.Println(postorderTraversal(root))
	fmt.Println(postorderTraversalIterative(root))
}
