package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time -> 0(N) Space -> 0(1)
func reOderLL(head *ListNode) *ListNode {
	count := 0
	tmp := head

	for tmp != nil {
		count++
		tmp = tmp.Next
	}

	mid := count / 2
	reverse := head
	var prev *ListNode
	for mid >= 0 {
		prev = reverse
		reverse = reverse.Next
		mid--
	}

	prev.Next = nil

	var tail *ListNode
	for reverse != nil {
		tmp = reverse.Next
		reverse.Next = tail
		tail = reverse
		reverse = tmp
	}

	curr := head
	for tail != nil {
		headTmp := curr.Next
		tailTmp := tail.Next
		curr.Next = tail
		tail.Next = headTmp
		curr = headTmp
		tail = tailTmp
	}

	return head
}

func main() {
	n0 := &ListNode{Val: 4}
	n1 := &ListNode{Val: 3, Next: n0}
	n2 := &ListNode{Val: 2, Next: n1}
	n3 := &ListNode{Val: 1, Next: n2}

	val := reOderLL(n3)

	for val != nil {
		fmt.Println(val.Val)
		val = val.Next
	}
}
