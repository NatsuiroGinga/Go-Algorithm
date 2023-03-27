package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	mid := (low + high) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = dfs(nums, low, mid-1)
	node.Right = dfs(nums, mid+1, high)
	return node
}
