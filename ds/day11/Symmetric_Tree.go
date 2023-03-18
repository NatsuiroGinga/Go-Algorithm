package day11

/*
First: check if `root` is nil, if it is, return true
Second: check if `root.left' and `root.right` are nil, means it is a leaf node, return true
Third: check if `root.left` and `root.right` are not both nil, and their values are equal, if not, return false
Fourth: check if `left.left` and `right.right` are mirror images, and `left.right` and `right.left` are mirror images
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	//return IsMirror(root.Left, root.Right)
	return isMirror(root)
}

func IsMirror(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && IsMirror(left.Left, right.Right) && IsMirror(left.Right, right.Left)
}

func isMirror(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)

	if root.Left != nil && root.Right == nil {
		return false
	} else if root.Left == nil && root.Right != nil {
		return false
	} else if root.Left != nil && root.Right != nil {
		stack = append(stack, root.Left)
		stack = append(stack, root.Right)
	}

	for len(stack) > 0 {
		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		right := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack)%2 != 0 || left.Val != right.Val {
			return false
		}

		if left.Left != nil && right.Right == nil {
			return false
		} else if left.Left == nil && right.Right != nil {
			return false
		} else if left.Left != nil && right.Right != nil {
			stack = append(stack, left.Left)
			stack = append(stack, right.Right)
		}

		if left.Right != nil && right.Left == nil {
			return false
		} else if left.Right == nil && right.Left != nil {
			return false
		} else if left.Right != nil && right.Left != nil {
			stack = append(stack, left.Right)
			stack = append(stack, right.Left)
		}
	}

	return true
}
