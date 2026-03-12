package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(log(N))
func isSymmetric(root TreeNode) bool {
	return recurse(root.Left, root.Right)
}

func recurse(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return recurse(p.Left, q.Right) && recurse(p.Right, q.Left)
}
