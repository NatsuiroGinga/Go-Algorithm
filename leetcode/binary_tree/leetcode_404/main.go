package main

import bt "Algorithm/leetcode/binary_tree"

var ans int

func sumOfLeftLeaves(root *bt.TreeNode) int {
	// 中序
	ans = 0
	inorder(root)
	return ans
}

func inorder(node *bt.TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		ans += node.Left.Val
	}

	inorder(node.Left)
	inorder(node.Right)
}
