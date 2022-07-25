/*
Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
k is a positive integer and is less than or equal to the length of the linked list. 
If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
You may not alter the values in the list's nodes, only nodes themselves may be changed.
*/

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	m := 0       // max length so far
	x, y := 0, 0 // index of max length

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true

		if i+1 < n && s[i] == s[i+1] {
			dp[i][i+1] = true
			m = 2
			x, y = i, i+1
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if j-i+1 >= m {
					m = j - i + 1
					x, y = i, j
				}
			}
		}
	}
	return s[x : y+1]
}
