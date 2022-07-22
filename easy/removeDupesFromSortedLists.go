/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. 
Return the linked list sorted as well.
*/

func deleteDuplicates(head *ListNode) *ListNode {
	if (head == nil) {return head}
	var cur *ListNode = head
	
	for cur.Next != nil {
			if(cur.Next.Val == cur.Val){
					cur.Next = cur.Next.Next;
			}else{cur = cur.Next}
	}
	
	return head
}
