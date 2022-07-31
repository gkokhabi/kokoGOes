/*
Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
The integer division should truncate toward zero, which means losing its fractional part. 
For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. 
For this problem, if the quotient is strictly greater than 231 - 1, 
then return 231 - 1, and if the quotient is strictly less than -231, then return -231.
*/
//aproach: d and c
//get the absoolute value
func abs(x int) int {
	if x < 0 {
			return -x
	}
	return x
}

func divide(dividend int, divisor int) int {
	if dividend == (-2<<30) && divisor == -1 {
			return (2<<30) - 1
	}
	
	dividendAbs := abs(dividend)
	divisorAbs := abs(divisor)
	res := 0
	
	for dividendAbs - divisorAbs  >= 0 {
			exponent := 0
			for dividendAbs - divisorAbs << exponent >= 0 {
					exponent++
			}
			res += 1 << (exponent - 1)
			dividendAbs -= divisorAbs << (exponent - 1)
	}
	if (divisor > 0) == (dividend >= 0) {
			return res
	}
	return -res
}

