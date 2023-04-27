package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}
	left := hasPathSum(root.Left, targetSum-root.Val)
	right := hasPathSum(root.Right, targetSum-root.Val)
	return left || right
}
