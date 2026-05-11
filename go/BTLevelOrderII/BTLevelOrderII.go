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
		currQueue := activeQueue[0]
		activeQueue = activeQueue[1:]
		nextQueue := make([]*TreeNode, 0)
		currResult := make([]int, 0)

		for len(currQueue) > 0 {
			node := currQueue[0]
			currQueue = currQueue[1:]

			if node == nil {
				continue
			}

			currResult = append(currResult, node.Val)

			nextQueue = append(nextQueue, node.Left)
			nextQueue = append(nextQueue, node.Right)
		}

		if len(nextQueue) > 0 {
			activeQueue = append(activeQueue, nextQueue)
		}

		if len(currResult) > 0 {
			result = append(result, currResult)
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func main() {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: left, Right: right}
	fmt.Println(levelOrder(root))
}
