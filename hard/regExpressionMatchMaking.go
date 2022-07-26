/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:
'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).
*/

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}
	prev := lists[0]
	var res *ListNode
	for i := 1; i < len(lists); i++ {
		res = mergeLists(prev, lists[i])
        prev= res      
	}

	return res
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	var head, temp *ListNode
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		head, temp = list1, list1
		list1 = list1.Next
	} else {
		head, temp = list2, list2
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			temp.Next = list1
			temp = temp.Next
			list1 = list1.Next
		} else {
			temp.Next = list2
			temp = temp.Next
			list2 = list2.Next
		}
	}

	if list1 != nil {
		temp.Next = list1
	}
	if list2 != nil {
		temp.Next = list2
	}

	return head
}
