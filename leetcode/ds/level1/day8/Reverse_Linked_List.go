package day8

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return dfs(head, nil)
}

// dfs cur: 当前节点 pre: 前一个节点
func dfs(cur, pre *ListNode) *ListNode {
	if cur == nil { // 递归终止条件
		return pre // 返回前一个节点
	}
	next := cur.Next
	// 反转
	cur.Next = pre
	// 递归, 当前节点变为前一个节点, 下一个节点变为当前节点
	return dfs(next, cur)
}
