package main

import "fmt"

type BinaryNode struct {
	Val   int
	Left  *BinaryNode
	Right *BinaryNode
}

// Time -> 0(N) Space -> 0(logN)
func countNodes(root *BinaryNode) int {
	count := 0
	recurse(root, &count)
	return count
}

func recurse(root *BinaryNode, count *int) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftSum, leftSize := recurse(root.Left, count)
	rightSum, rightSize := recurse(root.Right, count)

	currSum := root.Val + leftSum + rightSum
	currSize := leftSize + rightSize + 1

	if (currSum / currSize) == root.Val {
		*count++
	}

	return currSum, currSize
}

func main() {
	leaf1 := &BinaryNode{Val: 0, Left: nil, Right: nil}
	leaf2 := &BinaryNode{Val: 1, Left: nil, Right: nil}
	leaf3 := &BinaryNode{Val: 6, Left: nil, Right: nil}
	mid1 := &BinaryNode{Val: 8, Left: leaf1, Right: leaf2}
	mid2 := &BinaryNode{Val: 5, Left: nil, Right: leaf3}
	root := BinaryNode{Val: 4, Left: mid1, Right: mid2}
	fmt.Println(countNodes(&root))
}
