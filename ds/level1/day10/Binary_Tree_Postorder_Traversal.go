package day10

func postorderTraversal(root *TreeNode) []int {
	var (
		ans   []int
		stack []*TreeNode
	)

	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			// root.Val 加到 ans 的头部
			ans = append([]int{root.Val}, ans...)
			root = root.Right
		} else {
			root = stack[len(stack)-1].Left
			stack = stack[:len(stack)-1]
		}
	}

	return ans
}
