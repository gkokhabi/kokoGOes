/*
Write an algorithm to determine if a number n is happy.
A happy number is a number defined by the following process:
Starting with any positive integer, replace the number by the sum of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.
*/

func isHappy(n int) bool {
	visited := make(map[int]struct{})
	return helper(n, visited)
}

func helper(n int, visited map[int]struct{}) bool {
	s := strconv.Itoa(n)
	sum := 0
	for _, v := range s {
			tmp, _ := strconv.Atoi(string(v))
			sum += tmp * tmp
	}
	if sum == 1 { return true }
	if _, ok := visited[sum]; ok {
			return false
	} else {
			visited[sum] = struct{}{}
	}
	return helper(sum, visited)
}
