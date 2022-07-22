/*
Given an array nums of size n, return the majority element.
The majority element is the element that appears more than ⌊n / 2⌋ times. 
You may assume that the majority element always exists in the array.
*/

func majorityElement(nums []int) int {
	n := len(nums) / 2
	count := make(map[int]int)
	for _, v := range nums {
			count[v]++
			if count[v] > n { return v }
	}
	return -1
}

