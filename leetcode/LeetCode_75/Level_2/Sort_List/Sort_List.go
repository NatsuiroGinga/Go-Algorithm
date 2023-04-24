package Sort_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	var mid *ListNode
	for fast != nil && fast.Next != nil {
		mid = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid.Next = nil

	list1 := sortList(head)
	list2 := sortList(slow)

	return merge(list1, list2)
}

func merge(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := &ListNode{}
	head.Next = nil
	rear := head
	p1, p2 := list1, list2

	for ; p1 != nil && p2 != nil; rear = rear.Next {
		if p1.Val < p2.Val {
			rear.Next = p1
			p1 = p1.Next
		} else if p1.Val >= p2.Val {
			rear.Next = p2
			p2 = p2.Next
		}
	}

	if p1 != nil {
		rear.Next = p1
	}
	if p2 != nil {
		rear.Next = p2
	}

	return head.Next
}
