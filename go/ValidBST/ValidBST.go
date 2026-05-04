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

func isValidBST(root *TreeNode) bool {
	return checkBST(root, math.MinInt, math.MaxInt)
}

// Time -> 0(N) Space -> 0(logN)
func checkBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return checkBST(root.Left, min, root.Val) && checkBST(root.Right, root.Val, max)
}

func main() {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: left, Right: right}

	fmt.Println(isValidBST(root))
}
