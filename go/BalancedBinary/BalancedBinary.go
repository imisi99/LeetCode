package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(log(N))
func isBalanced(root *TreeNode) bool {
	_, balanced := check(root)
	return balanced
}

func check(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	height1, balanced1 := check(root.Left)
	height2, balanced2 := check(root.Right)

	balanced := balanced1 && balanced2

	diff := height1 - height2
	if diff < -1 || diff > 1 {
		balanced = false
	} else {
		balanced = balanced && true
	}

	height := max(height1, height2) + 1

	return height, balanced
}
