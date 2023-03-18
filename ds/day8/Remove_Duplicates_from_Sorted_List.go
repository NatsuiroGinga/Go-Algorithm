package day8

/*
Intuition: tow pointers, left and right, left is the last node that is not duplicate, right is the current node.
Approach: if left.Val and right.Val are not equal, move left to right, else move right to right.Next, and remove the duplicate node until left.Val and right.Val are not equal or right is nil.
Time Complexity: O(N)
Space Complexity: O(1)
*/
func deleteDuplicates(head *ListNode) *ListNode {
	// init left and right pointers, go through the list
	for left, right := head, head; right != nil; right = right.Next {
		// while left.Val and right.Val are equal and left and right are not equal
		for left != right && left.Val == right.Val {
			// left.Next points to right.Next, remove the duplicate node
			left.Next = right.Next
			// move right to right.Next
			right = right.Next
			// if right is nil, means the list is finished, return head
			if right == nil {
				return head
			}
		}
		// if left and right are not equal, move left to right
		left = right
	}

	return head
}
