package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N) Space -> 0(N)
func kthSmallest(root *TreeNode, k int) int {
	array := make([]int, 0)
	recurse(root, &array)
	return array[k-1]
}

func recurse(root *TreeNode, array *[]int) {
	if root == nil {
		return
	}

	recurse(root.Left, array)
	*array = append(*array, root.Val)
	recurse(root.Right, array)
}

// Time -> 0 max(logN, k) Space -> 0(logN)
func kthSmallestI(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for k > 1 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--

		if node.Right != nil {
			stack = append(stack, node.Right)
			curr := node.Right.Left
			for curr != nil {
				stack = append(stack, curr)
				curr = curr.Left
			}
		}
	}

	return stack[len(stack)-1].Val
}

func main() {
	lright := &TreeNode{2, nil, nil}
	left := &TreeNode{1, nil, lright}
	right := &TreeNode{4, nil, nil}
	root := &TreeNode{3, left, right}

	fmt.Println(kthSmallest(root, 3))
	fmt.Println(kthSmallestI(root, 3))
}
