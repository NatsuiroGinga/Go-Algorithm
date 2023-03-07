package day3

// recursive
/*func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}*/

// iterative
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head, rear *ListNode

	// if one of the lists is empty, return the other list
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// find head node of merged list by comparing the first node of each list
	if list1.Val <= list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	rear = head // rear is the last node of merged list

	// compare the first node of each list and add the smaller one to the merged list
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			rear.Next = list1  // add list1 to the merged list
			list1 = list1.Next // move list1 to the next node
		} else {
			rear.Next = list2  // add list2 to the merged list
			list2 = list2.Next // move list2 to the next node
		}

		rear = rear.Next // move rear to the next node
	}

	// add the remaining nodes of list1 or list2 to the merged list
	if list1 == nil { // list1 is empty
		rear.Next = list2 // add list2 to the merged list
	} else if list2 == nil { // list2 is empty
		rear.Next = list1 // add list1 to the merged list
	}

	return head
}
