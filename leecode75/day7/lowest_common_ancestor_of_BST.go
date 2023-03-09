package day7

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root
	for {
		if root.Val < p.Val && root.Val < q.Val {
			cur = cur.Right
		} else if root.Val > p.Val && root.Val > q.Val {
			cur = cur.Left
		} else {
			return cur
		}
	}
}
