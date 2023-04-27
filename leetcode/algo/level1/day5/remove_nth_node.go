package day5

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		pre  = &ListNode{Next: head}
		slow = pre
		fast = pre
	)

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return pre.Next
}
