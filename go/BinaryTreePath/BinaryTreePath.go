package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	recurse(root, "", &result)
	return result
}

// Time -> NlogN Space -> NlogN
func recurse(root *TreeNode, path string, result *[]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*result = append(*result, path+fmt.Sprintf("%v", root.Val))
		return
	}

	path += fmt.Sprintf("%v", root.Val) + "->"

	recurse(root.Left, path, result)
	recurse(root.Right, path, result)
}
