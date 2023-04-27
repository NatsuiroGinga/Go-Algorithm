package leetcode_98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode = nil

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := isValidBST(root.Left)

	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root

	right := isValidBST(root.Right)

	return left && right
}
