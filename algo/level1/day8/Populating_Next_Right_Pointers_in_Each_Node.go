package day8

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	//dfs(root, nil)
	bfs(root)
	return root
}

func dfs(cur, next *Node) {
	if cur == nil {
		return
	}

	cur.Next = next

	dfs(cur.Left, cur.Right)

	if cur.Next != nil {
		dfs(cur.Right, cur.Next.Left)
	} else {
		dfs(cur.Right, nil)
	}
}

func bfs(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		var rightNode *Node

		for i := len(queue) - 1; i >= 0; i-- {
			cur := queue[0]
			queue = queue[1:]
			cur.Next = rightNode
			rightNode = cur

			if cur.Right != nil {
				queue = append(queue, cur.Right)
				queue = append(queue, cur.Left)
			}
		}
	}

	return root
}
