package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Time -> 0(N) Space -> 0(N)
func populate(tree *Node) *Node {
	queue := make([][]*Node, 0)
	queue = append(queue, []*Node{tree})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		next := make([]*Node, 0)
		for len(node) > 0 {
			curr := node[0]
			node = node[1:]

			if curr == nil {
				continue
			}

			if len(node) > 0 {
				curr.Next = node[0]
			}

			next = append(next, curr.Left)
			next = append(next, curr.Right)
		}

		if len(next) > 0 {
			queue = append(queue, next)
		}
	}
	return tree
}
