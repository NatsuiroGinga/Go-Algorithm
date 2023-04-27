package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	rear := head
	sum := 0

	for l1 != nil || l2 != nil {
		sum /= 10
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		rear.Next = &ListNode{Val: sum % 10}
		rear = rear.Next
	}

	if sum/10 == 1 {
		rear.Next = &ListNode{Val: 1}
	}

	return head.Next
}
