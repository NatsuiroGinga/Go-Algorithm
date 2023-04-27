package leetcode_501

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre *TreeNode
var count, maxCount int
var ans []int

func findMode(root *TreeNode) []int {
	ans = make([]int, 0)
	count, maxCount = 1, 1
	pre = nil
	inorder(root)
	return ans
}

func inorder(cur *TreeNode) {
	if cur == nil {
		return
	}

	inorder(cur.Left)

	if pre != nil && pre.Val == cur.Val {
		count++
	} else {
		count = 1
	}
	pre = cur

	if count > maxCount {
		maxCount = count
		ans = ans[:0]
		ans = append(ans, cur.Val)
	} else if count == maxCount {
		ans = append(ans, cur.Val)
	}

	inorder(cur.Right)
}
