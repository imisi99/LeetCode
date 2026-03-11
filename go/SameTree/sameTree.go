package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(log(N))
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return recurse(p, q)
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

	return recurse(p.Left, q.Left) && recurse(p.Right, q.Right)
}
