package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	mp := make(map[int]int)
	mp[0] = 1
	return dfs(root, 0, targetSum, mp)
}

func dfs(root *TreeNode, curSum int, targetSum int, mp map[int]int) int {
	if root == nil {
		return 0
	}

	curSum += root.Val
	res := mp[root.Val-targetSum]
	mp[curSum]++
	res += dfs(root.Left, curSum, targetSum, mp) + dfs(root.Right, curSum, targetSum, mp)
	mp[curSum]--

	return res
}
