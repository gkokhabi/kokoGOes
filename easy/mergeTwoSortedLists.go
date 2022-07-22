/*
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. 
The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	//0 node
	if list1 == nil || list2 == nil {
		if list1 != nil {
			return list1
		} else if list2 != nil {
			return list2
		}
		return nil
	}

	//over 1 nodes
	var root *ListNode
	var prev *ListNode
	//set root node
	if list1.Val <= list2.Val {
		root = list1
		list1 = list1.Next
	} else {
		root = list2
		list2 = list2.Next
	}
	prev = root

	//set node
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			prev.Next = list1
			prev = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			prev = list2
			list2 = list2.Next
		}
	}

	//set last node
	if list1 == nil {
		prev.Next = list2
	} else {
		prev.Next = list1
	}

	return root
}
