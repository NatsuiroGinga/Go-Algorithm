package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	/*if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root*/
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur := root

	for {
		if cur.Val < val {
			if cur.Right != nil {
				cur = cur.Right
			} else {
				cur.Right = &TreeNode{Val: val}
				break
			}
		} else {
			if cur.Left != nil {
				cur = cur.Left
			} else {
				cur.Left = &TreeNode{Val: val}
				break
			}
		}
	}

	return root
}
