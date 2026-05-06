package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time -> 0(N)
func levelOrder(root *TreeNode) [][]int {
	activeQueue := make([][]*TreeNode, 0)
	result := make([][]int, 0)
	activeQueue = append(activeQueue, []*TreeNode{root})

	for len(activeQueue) > 0 {
		node := activeQueue[0]
		activeQueue = activeQueue[1:]
		nextLevel := make([]*TreeNode, 0)
		currLevel := make([]int, 0)

		for len(node) > 0 {
			currNode := node[0]
			node = node[1:]

			if currNode == nil {
				continue
			}
			currLevel = append(currLevel, currNode.Val)

			nextLevel = append(nextLevel, currNode.Left)
			nextLevel = append(nextLevel, currNode.Right)
		}

		if len(nextLevel) > 0 {
			activeQueue = append(activeQueue, nextLevel)
		}

		if len(currLevel) > 0 {
			result = append(result, currLevel)
		}

	}
	return result
}

func main() {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: left, Right: right}
	fmt.Println(levelOrder(root))
}
