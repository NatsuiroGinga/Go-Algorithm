package leetcode_530

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode
var ans int

func getMinimumDifference(root *TreeNode) int {
	pre = nil
	ans = math.MaxInt
	inorder(root)
	return ans
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	if pre != nil {
		ans = min(ans, node.Val-pre.Val)
	}
	pre = node
	inorder(node.Right)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
