/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.
*/

func romanToInt(s string) int {

	type Value struct {
		Value  uint16
		Before byte
	}

	dic := map[byte]Value{
		"I"[0]: {1, " "[0]}, "V"[0]: {5, "I"[0]}, "X"[0]: {10, "I"[0]}, "L"[0]: {50, "X"[0]}, "C"[0]: {100, "X"[0]}, "D"[0]: {500, "C"[0]}, "M"[0]: {1000, "C"[0]},
	}

	sum := uint16(0)
	before := " "[0]
	length := uint16(len(s))

	for i := uint16(0); i < length; i++ {

		current := s[i]
		d := dic[current]
		sum += d.Value

		if before == d.Before {
			sum -= dic[d.Before].Value * 2
		}

		before = current

	}
	return int(sum)
}

