package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	findPathSums(root, targetSum, []int{}, &result)
	return result
}

func findPathSums(root *TreeNode, targetSum int, path []int, result *[][]int) {
	if root == nil {
		return
	}

	newTarget := targetSum - root.Val
	path = append(path, root.Val)

	if newTarget == 0 && root.Left == nil && root.Right == nil {
		newPath := append([]int{}, path...)
		*result = append(*result, newPath)
		return
	}

	findPathSums(root.Left, newTarget, path, result)
	findPathSums(root.Right, newTarget, path, result)
}

func main() {
	left := &TreeNode{Val: 2}
	a := &TreeNode{Val: -1}
	right := &TreeNode{Val: 3, Left: a}
	root := &TreeNode{Val: 1, Left: left, Right: right}

	fmt.Println(pathSum(root, 3))
}
