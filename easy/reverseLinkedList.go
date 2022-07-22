/*
Given the head of a singly linked list, reverse the list, and return the reversed list.
*/

func reverseList(head *ListNode) *ListNode {
	_, root := impl(head)
	return root
}

func impl(head *ListNode) (cur, root *ListNode) {
	if head == nil {
			return nil, nil
	}
	
	prev, root := impl(head.Next)
	if prev == nil {
			return head, head
	}
	
	cur = &ListNode{Val:head.Val}
	prev.Next = cur
	return cur, root
}
