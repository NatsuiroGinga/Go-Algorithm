package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Values []int
	cur    int
}

func Constructor(root *TreeNode) BSTIterator {
	stack := list.New()
	var values []int

	for stack.Len() > 0 || root != nil {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		root = stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())

		values = append(values, root.Val)

		root = root.Right
	}

	return BSTIterator{values, -1}
}

func (iter *BSTIterator) Next() int {
	iter.cur++
	return iter.Values[iter.cur]
}

func (iter *BSTIterator) HasNext() bool {
	return iter.cur+1 < len(iter.Values)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
