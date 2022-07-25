/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
Merge all the linked-lists into one sorted linked-list and return it.
*/

func isMatch(s string, p string) bool {
	if p == "" { return s == ""}
	firstMatch := !(s == "") && (p[0] == s[0] || p[0] == 46)
	//first char in pattern equals to s[0] or .
	if len(p) >= 2 && p[1] == 42 {
		//f p[0] is followed by a *
		//if the first chat matches with p[0] then check its following substring || //if not, skip the c* in pattern
			return (firstMatch && isMatch(s[1:], p)) || isMatch(s,p[2:]) 
	}else{
		// if pattern length is 1 or p[0] is not followed by *
			return firstMatch && isMatch(s[1:],p[1:])
	}
}
