/*
Given an integer x, return true if x is palindrome integer.
An integer is a palindrome when it reads the same backward as forward.
For example, 121 is a palindrome while 123 is not.
*/

func isPalindrome(x int) bool {
    
	if x < 0 return false

	tmp := x
	var valueList []int
	times := 1

	for tmp != 0 {
		valueList = append(valueList, tmp%10)
		tmp /= 10
	}

	tmp = 0
	for i := len(valueList) - 1; i >= 0; i-- {
		tmp += valueList[i] * times
		times *= 10
	}

	return x == tmp
    
}

