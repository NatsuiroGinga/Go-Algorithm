package leetcode_654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return construct(nums, 0, len(nums)-1)
}

func construct(nums []int, low, high int) *TreeNode {
	if low > high {
		return nil
	}
	max, index := MaxNum(nums, low, high)

	return &TreeNode{
		Val:   max,
		Left:  construct(nums, 0, index-1),
		Right: construct(nums, index+1, high),
	}
}

func MaxNum(nums []int, from, to int) (num, index int) {
	num = nums[from]
	index = from
	for i := from + 1; i <= to; i++ {
		if nums[i] > num {
			num = nums[i]
			index = i
		}
	}
	return num, index
}
