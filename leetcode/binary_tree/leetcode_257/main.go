package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans []string

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	ans = make([]string, 0)
	preorder(root, "")
	return ans
}

func preorder(node *TreeNode, path string) {
	if node.Left == nil && node.Right == nil {
		ans = append(ans, path+strconv.Itoa(node.Val))
		return
	}

	if node.Left != nil {
		preorder(node.Left, path+strconv.Itoa(node.Val)+"->")
	}

	if node.Right != nil {
		preorder(node.Right, path+strconv.Itoa(node.Val)+"->")
	}
}
