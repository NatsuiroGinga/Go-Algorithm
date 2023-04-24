package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	yes, no := robTree(root)
	return max(yes, no)
}

func robTree(cur *TreeNode) (yes, no int) {
	if cur == nil {
		return 0, 0
	}

	leftYes, leftNo := robTree(cur.Left)
	rightYes, rightNo := robTree(cur.Right)

	// 偷当前节点, 则左右子树节点不能偷
	robCur := cur.Val + leftNo + rightNo
	// 不偷当前节点, 则左右子树节点可以偷
	notRobCur := max(leftYes, leftNo) + max(rightNo, rightYes)

	return robCur, notRobCur
}
