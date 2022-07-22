/*
You are given two non-empty linked lists representing two non-negative integers. 
The digits are stored in reverse order, and each of their nodes contains a single digit. 
Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
			n1, n2 := 0, 0
			if l1 != nil {
					n1, l1 = l1.Val, l1.Next
			}
			if l2 != nil {
					n2, l2 = l2.Val, l2.Next
			}
			num := n1 + n2 + carry
			carry = num / 10
			cur.Next = &ListNode{Val:num%10, Next:nil}
			cur = cur.Next
	}
	return head.Next
}