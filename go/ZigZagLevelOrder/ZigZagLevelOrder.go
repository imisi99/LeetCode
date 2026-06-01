package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := make([][]*TreeNode, 0)
	result := make([][]int, 0)
	queue = append(queue, []*TreeNode{root})
	dir := "left"

	for len(queue) != 0 {
		activeQueue := queue[0]
		queue = queue[1:]
		nextQueue := make([]*TreeNode, 0)
		currResult := make([]int, 0)

		for _, node := range activeQueue {
			if node == nil {
				continue
			}

			currResult = append(currResult, node.Val)

			nextQueue = append(nextQueue, node.Left)
			nextQueue = append(nextQueue, node.Right)
		}

		if len(currResult) > 0 {
			if dir == "right" {
				for i, j := 0, len(currResult)-1; i < j; i, j = i+1, j-1 {
					currResult[i], currResult[j] = currResult[j], currResult[i]
				}
				result = append(result, currResult)
				dir = "left"
			} else {
				result = append(result, currResult)
				dir = "right"
			}
		}

		if len(nextQueue) > 0 {
			queue = append(queue, nextQueue)
		}
	}
	return result
}

func main() {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: left, Right: right}
	fmt.Println(zigzagLevelOrder(root))
}
