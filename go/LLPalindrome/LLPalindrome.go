package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	return approachII(head)
}

// Time -> 0(N) Space -> 0(N)
func approachI(head *ListNode) bool {
	array := make([]int, 0)

	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	start, end := 0, len(array)-1

	for start < end {
		if array[start] != array[end] {
			return false
		}
		start++
		end--
	}

	return true
}

// Time -> 0(N) Space -> 0(1)
func approachII(head *ListNode) bool {
	count := 0
	curr := head
	for curr != nil {
		curr = curr.Next
		count++
	}

	mid := count / 2

	reverse := head
	for mid > 0 {
		reverse = reverse.Next
		mid--
	}

	if count%2 == 1 {
		reverse = reverse.Next
	}

	var endLL *ListNode
	for reverse != nil {
		tmp := reverse.Next
		reverse.Next = endLL
		endLL = reverse
		reverse = tmp
	}

	for endLL != nil {
		if endLL.Val != head.Val {
			return false
		}
		endLL = endLL.Next
		head = head.Next
	}

	return true
}
