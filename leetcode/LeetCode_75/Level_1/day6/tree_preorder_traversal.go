package day6

type Node struct {
	Val      int
	Children []*Node
}

// recursive
/*func preorder(root *Node) []int {
	var ret []int
	dfs(root, &ret)
	return ret
}

func dfs(root *Node, ret *[]int) {
	if root == nil {
		return
	}

	*ret = append(*ret, root.Val)

	for _, v := range root.Children {
		dfs(v, ret)
	}
}*/

// iterative
func preorder(root *Node) []int {
	var ret []int
	if root == nil {
		return ret
	}

	// stack
	stack := []*Node{root}

	// while stack is not empty
	for len(stack) > 0 {
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// push
		ret = append(ret, node.Val)
		// push children
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return ret
}
