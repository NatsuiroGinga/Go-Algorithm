package leetcode_24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		pre, pos := cur.Next, cur.Next.Next
		pos.Next = pre // 后指前
		cur.Next = pos // 虚拟结点 指 后
		cur = pos.Next //虚拟节点向后移动两个单位
		pre.Next = cur //前 指 虚拟节点
	}

	return dummyHead.Next
}
