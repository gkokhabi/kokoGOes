/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

func singleNumber(nums []int) int {
	n := 0
data := make(map[int]int)
for _, val := range nums {
	data[val] = data[val] + 1
}
for key, val := range data {
	if val == 1 {
		n = key
		break
	}
}
return n
}
