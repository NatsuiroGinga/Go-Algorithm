package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	pre := slow
	slow = slow.Next
	pre.Next = nil

	for slow != nil {
		tmp := slow.Next
		slow.Next = pre
		pre = slow
		slow = tmp
	}

	slow = pre
	fast = head
	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}

	return true
}

func main() {

}
