package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(log(N))
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findDepth(root, 1)
}

func findDepth(root *TreeNode, depth int) int {
	if root.Left == nil && root.Right == nil {
		return depth
	}

	newDepth := depth + 1
	leftDepth := math.MaxInt
	rightDepth := math.MaxInt
	if root.Left != nil {
		leftDepth = findDepth(root.Left, newDepth)
	}
	if root.Right != nil {
		rightDepth = findDepth(root.Right, newDepth)
	}

	return min(leftDepth, rightDepth)
}

func main() {
	left := &TreeNode{Val: 9}
	// right := &TreeNode{Val: 20}
	root := &TreeNode{Val: 3, Left: left}
	fmt.Println(minDepth(root))
}
