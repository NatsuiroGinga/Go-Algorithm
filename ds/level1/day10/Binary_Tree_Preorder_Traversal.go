package day10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		ans   []int
		stack []*TreeNode
	)

	stack = append(stack, root)

	for len(stack) > 0 {
		if root == nil && len(stack) > 0 {
			parent := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = parent.Right
		} else {
			ans = append(ans, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
	}

	return ans
}
