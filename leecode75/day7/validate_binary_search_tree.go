package day7

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var (
		stack []*TreeNode // 用于存储节点
		pre   *TreeNode   // 用于存储前一个节点
	)

	if root == nil {
		return false
	}

	for root != nil || len(stack) > 0 { // 中序遍历
		for root != nil { // 将左子树全部入栈
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1] // 取出栈顶元素
		stack = stack[:len(stack)-1]

		if pre != nil && root.Val <= pre.Val { // 判断是否满足二叉搜索树的条件
			return false
		}
		// 更新前一个节点
		pre = root
		// 更新当前节点 为当前节点的右子树
		root = root.Right
	}

	return true
}
