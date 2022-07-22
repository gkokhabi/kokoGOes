/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) return false

	records := [26]int{}
	for _,v := range s {
			records[v - 'a']++
	}
	for _,v := range t {
			records[v - 'a']--
	}
	for _,v := range records {
			if v != 0 {
					return false
			}
	}
	return true
	
}

