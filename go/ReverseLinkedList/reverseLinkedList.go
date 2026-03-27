package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	if head == nil {
		return nil
	}
	for head.Next != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	head.Next = prev
	return head
}

func main() {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2, Next: node1}
	root := &ListNode{Val: 1, Next: node2}
	head := reverseList(root)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
