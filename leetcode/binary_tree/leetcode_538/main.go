package leetcode_538

var pre *TreeNode

func convertBST(root *TreeNode) *TreeNode {
	pre = nil
	convert(root)
	return root
}

func convert(cur *TreeNode) {
	if cur == nil {
		return
	}
	convert(cur.Right)

	if pre != nil {
		cur.Val += pre.Val
	}
	pre = cur

	convert(cur.Left)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
