/*
Given a string s, find the first non-repeating character in it and return its index. 
If it does not exist, return -1.
*/


func firstUniqChar(s string) int {
	tmp := make([]int, 26)
	for _, v := range s {
			tmp[v - 'a']++
	}
	for i, v := range s {
			if tmp[v - 'a'] == 1 { return i }
	}
	return -1
}
