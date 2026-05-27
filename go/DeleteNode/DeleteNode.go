package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func deleteNode(node *ListNode) {
	var prev *ListNode
	for node.Next != nil {
		node.Val = node.Next.Val
		prev = node
		node = node.Next
	}
	prev.Next = nil
}

func main() {
	two := &ListNode{Val: 3}
	one := &ListNode{Val: 2, Next: two}
	head := &ListNode{Val: 1, Next: one}

	deleteNode(head)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
