package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(log(N))
func maxDepth(root *TreeNode) int {
	return findDepth(root, 0)
}

func findDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	newDepth := depth + 1
	leftDepth := findDepth(root.Left, newDepth)
	rightDepth := findDepth(root.Right, newDepth)

	return max(leftDepth, rightDepth)
}

func main() {
	left := &TreeNode{Val: 9}
	right := &TreeNode{Val: 20}
	root := &TreeNode{Val: 3, Left: left, Right: right}
	fmt.Println(maxDepth(root))
}
