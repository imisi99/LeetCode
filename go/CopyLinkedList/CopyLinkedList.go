package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Time -> 0(N) Space -> 0(N)
func copyNode(head *Node) *Node {
	nodeMap := make(map[*Node]*Node, 0)

	var newHead, start *Node
	actualHead := head

	for head != nil {
		if start == nil {
			start = &Node{}
			newHead = start
		}

		start.Val = head.Val
		nodeMap[head] = start

		if head.Next != nil {
			start.Next = &Node{}
		}

		head = head.Next
		start = start.Next
	}

	head = newHead

	for newHead != nil && actualHead != nil {
		newHead.Random = nodeMap[actualHead.Random]
		newHead = newHead.Next
		actualHead = actualHead.Next
	}

	return head
}
