package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Array []int
	CurrIdx int
}

func Constructor(root *TreeNode) *BSTIterator {
	array := make([]int, 0)

	// Time -> 0(N) Space -> 0(N)
	func InOrder(root *TreeNode, array *[]int) {
		if root == nil {
			return 
		}

		InOrder(root.Left, array)
		*array = append(*array, root.Val)
		InOrder(root.Right, array)
	}

	InOrder(root, &array)

	return &BSTIterator{Array: array, CurrIdx: -1}
}

// Time -> 0(1) Space -> 0(1)
func (this *BSTIterator) Next() int {
	this.CurrIdx++
	return this.Array[this.CurrIdx-1]
}

// Time -> 0(1) Space -> 0(1)
func (this *BSTIterator) HasNext() bool {
	return !len(this.Array) == this.CurrIdx
}

type BSTIteratorI struct {
	Array []*TreeNode
}

func ConstructorI(root *TreeNode) *BSTIteratorI {
	iterator := &BSTIteratorI{Array: []*TreeNode}
	iterator.LeftMost(root)
	return iterator
}

// Time -> 0(h) Space -> 0(h)
func (this *BSTIteratorI) LeftMost(root *TreeNode) {
	for root != nil {
		this.Array = append(this.Array, root.Val)
		root = root.Left
	}
}

// Time -> 0(h) Space -> 0(h) on avg 0(1)
func (this *BSTIteratorI) Next() int {
	curr := this.Array[len(this.Array)-1]
	this.Array = this[:len(this.Array)-1]
	this.LeftMost(curr.Right)
	return curr.Val
}

// Time -> 0(h) Space -> 0(1)
func (this *BSTIteratorI) HasNext() bool {
	return len(this.Array) > 0
}