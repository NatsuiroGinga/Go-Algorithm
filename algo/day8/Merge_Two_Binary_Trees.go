package day8

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		root := &TreeNode{Val: root1.Val + root2.Val}
		root.Left = mergeTrees(root1.Left, root2.Left)
		root.Right = mergeTrees(root1.Right, root2.Right)
		return root
	}
	if root1 != nil {
		return root1
	}
	return root2
}
