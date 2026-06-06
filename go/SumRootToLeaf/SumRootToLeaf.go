package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	traverse(root, "", &sum)
	return sum
}

func traverse(root *TreeNode, path string, sum *int) {
	if root == nil {
		return
	}

	path += fmt.Sprintf("%v", root.Val)
	if root.Left == nil && root.Right == nil {
		pathInt, _ := strconv.Atoi(path)
		*sum += pathInt
		return
	}

	traverse(root.Left, path, sum)
	traverse(root.Right, path, sum)
}

func main() {
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 1, Left: left, Right: right}
	fmt.Println(sumNumbers(root))
}
