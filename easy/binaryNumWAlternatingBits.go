/*
Given a positive integer, check whether it has alternating bits: namely, 
if two adjacent bits will always have different values.
*/

func hasAlternatingBits(n int) bool {
    
	for n > 0 {
        if n%2 == (n>>1)&1 { return false }
		n = n >> 1
	}
    
	return true
    
}
