package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var mp map[int]struct{}
	return find(root, k, mp)
}

func find(root *TreeNode, k int, mp map[int]struct{}) bool {
	if root == nil {
		return false
	}
	if _, ok := mp[k-root.Val]; ok {
		return true
	}
	mp[root.Val] = struct{}{}
	return find(root.Left, k, mp) || find(root.Right, k, mp)
}
