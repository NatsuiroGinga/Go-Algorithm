package day3

func reverseList(head *ListNode) *ListNode {
	var p, next, pre *ListNode
	p = head
	pre = nil

	for p != nil {
		next = p.Next
		p.Next = pre
		pre = p
		p = next
	}

	return pre
}
