package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	recursiveSwap(root)
	return root
}

// Time -> 0(N) Space -> 0(log(N))
func recursiveSwap(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	recursiveSwap(root.Left)
	recursiveSwap(root.Right)
}
