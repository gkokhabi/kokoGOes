/*
You are a product manager and currently leading a team to develop a new product. 
Unfortunately, the latest version of your product fails the quality check. 
Since each version is developed based on the previous version, all the versions after a bad version are also bad.
Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. 
You should minimize the number of calls to the API.
*/

func firstBadVersion(n int) int {
	if n == 1 return 1

	var isMidBad bool
	for s, e, mid := 1, n, (1 + n) / 2; mid >= 1; mid = (s+e)/2 {
			isMidBad = isBadVersion(mid)
			
			if isMidBad && !isBadVersion(mid-1) {
					return mid
			} else if isMidBad {
					e = mid - 1
			} else {
					s = mid + 1
			}
	}
	
	return 1
}
