/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. 
(i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/


func isSubsequence(s string, t string) bool {
	if len(s) == 0 return true

	indexS, indexT := 0, 0
	for indexT < len(t) {
		if t[indexT] == s[indexS] {
			indexS++
			if indexS == len(s) {
				return true
			}
		}
		indexT++
	}

	return false
}
