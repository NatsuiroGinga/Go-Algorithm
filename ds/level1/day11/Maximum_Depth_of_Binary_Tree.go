package day11

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
First: check if `root` is nil, if it is, return 0
Second: get the max depth of left and right subtree
Third: return the max depth of left and right subtree plus 1
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return MaxInt(left, right) + 1
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
