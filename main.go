package main

func main() {

}

func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	if root == nil {
		return 0
	}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if k == 1 {
			return root.Val
		}
		k--

		root = root.Right
	}

	return 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
