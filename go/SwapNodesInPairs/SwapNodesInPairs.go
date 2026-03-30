package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func swapPairs(head *ListNode) *ListNode {
	swap := head
	if head != nil && head.Next != nil {
		swap = head.Next
	}
	var prev *ListNode
	for head != nil && head.Next != nil {
		next := head.Next.Next
		tmpHead := head
		head = head.Next
		if prev != nil {
			prev.Next = head
		}
		head.Next = tmpHead
		head.Next.Next = next
		prev = head.Next
		head = next
	}
	return swap
}

func main() {
	g := &ListNode{Val: 5}
	f := &ListNode{Val: 4, Next: g}
	d := &ListNode{Val: 3, Next: f}
	s := &ListNode{Val: 2, Next: d}
	a := &ListNode{Val: 1, Next: s}

	head := swapPairs(a)

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
