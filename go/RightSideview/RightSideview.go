package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(logN)
func rightSideView(root *TreeNode) []int {
	queue := make([][]*TreeNode, 0)
	queue = append(queue, []*TreeNode{root})
	visible := make([]int, 0)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		next := make([]*TreeNode, 0)

		for len(curr) > 0 {
			node := curr[0]
			curr = curr[1:]

			if node.Left != nil {
				next = append(next, node.Left)
			}

			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		if len(next) > 0 {
			queue = append(queue, next)
			visible = append(visible, next[len(next)-1].Val)
		}

	}
	return visible
}
