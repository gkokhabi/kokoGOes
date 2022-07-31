/*
Given a linked list, swap every two adjacent nodes and return its head. 
You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)
*/


func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { 
		return head 
	}
	tmp := head.Next
	head.Next = swapPairs(head.Next.Next)
	tmp.Next = head
	return tmp
}
