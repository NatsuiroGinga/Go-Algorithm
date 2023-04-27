package leetcode_108

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return recursion(nums, 0, len(nums)-1)
}

func recursion(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	mid := (low + high) / 2

	return &TreeNode{
		Val:   nums[mid],
		Left:  recursion(nums, low, mid-1),
		Right: recursion(nums, mid+1, high),
	}
}
