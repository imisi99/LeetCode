package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(logN)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	newTarget := targetSum - root.Val
	if newTarget == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	left := hasPathSum(root.Left, newTarget)
	right := hasPathSum(root.Right, newTarget)

	return left || right
}

func main() {
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 1, Left: left, Right: right}

	fmt.Println(hasPathSum(root, 3))
}
