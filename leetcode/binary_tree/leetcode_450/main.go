package leetcode_450

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		}

		min := root.Right
		for min.Left != nil {
			min = min.Left
		}
		root.Val = min.Val
		root.Right = deleteNode(root.Right, root.Val)
	}

	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
