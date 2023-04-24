package day6

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
/*func levelOrder(root *TreeNode) [][]int {
	var (
		ret   [][]int
		queue []*TreeNode
	)

	if root == nil {
		return ret
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		var subList []int

		for i := 0; i < size; i++ {
			// pop
			node := queue[0]
			queue = queue[1:]
			// add left child
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// add right child
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			subList = append(subList, node.Val)
		}
		// add sublist to ret
		ret = append(ret, subList)
	}

	return ret
}*/

// dfs
func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	dfs(root, 0, &ret)
	return ret
}

func dfs(root *TreeNode, height int, res *[][]int) {
	// recursion terminator
	if root == nil {
		return
	}
	// check if the height is equal to the length of the result
	if height == len(*res) {
		// append a new list
		*res = append(*res, []int{})
	}
	// add the current node.val to the res
	(*res)[height] = append((*res)[height], root.Val)
	// process the left and right child
	dfs(root.Left, height+1, res)
	dfs(root.Right, height+1, res)
}
