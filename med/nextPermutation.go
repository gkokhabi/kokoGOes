/*
A permutation of an array of integers is an arrangement of its members into a sequence or linear order.

For example, for arr = [1,2,3], the following are considered permutations of arr: [1,2,3], [1,3,2], [3,1,2], [2,3,1].
The next permutation of an array of integers is the next lexicographically greater permutation of its integer. 
More formally, if all the permutations of the array are sorted in one container according to their lexicographical order, 
then the next permutation of that array is the permutation that follows it in the sorted container. 
If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).

For example, the next permutation of arr = [1,2,3] is [1,3,2].
Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
Given an array of integers nums, find the next permutation of nums.

The replacement must be in place and use only constant extra memory.
*/

func nextPermutation(nums []int) {
	//get the index of last peak.
	lastPeakIdx := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			lastPeakIdx = i
			break
		}
	}
	//if the given input is in descending order.
	if lastPeakIdx == -1 {
		sort.Ints(nums)
		return
	}
	//if the given input is not in descending order.
	idx := lastPeakIdx
	for i := lastPeakIdx; i < len(nums); i++ {
		if nums[i] > nums[lastPeakIdx-1] && nums[i] < nums[idx] {
			//find an smallest element which exists between nums[lastPeakIdx-1] and nums [lastPeakIdx]
			idx = i
		}
	}
	nums[idx], nums[lastPeakIdx-1] = nums[lastPeakIdx-1], nums[idx]
	sort.Ints(nums[lastPeakIdx:])
}

