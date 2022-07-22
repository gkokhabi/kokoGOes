/*
Given an array of integers nums which is sorted in ascending order, and an integer target, 
write a function to search target in nums. If target exists, then return its index. 
Otherwise, return -1.
You must write an algorithm with O(log n) runtime complexity.
*/

func search(nums []int, target int) int {
	middle, left, right := 0, 0, len(nums)-1

for left <= right {
	middle = left + (right-left)/2

	if nums[middle] == target {
		return middle
	}

	if nums[middle] < target {
		left = middle + 1
	} else {
		right = middle - 1
	}
}

return -1
}
