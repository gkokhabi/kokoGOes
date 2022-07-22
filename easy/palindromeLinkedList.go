/*
Given the head of a singly linked list, return true if it is a palindrome.
*/


func isPalindrome(head *ListNode) bool {
	values := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		values = append(values, node.Val)
	}
	for i := 0; i < len(values) / 2; i++ {
		if values[i] != values[len(values)-1-i] {
			return false
		}
	}
	return true
}
