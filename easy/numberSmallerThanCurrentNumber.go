/*
Given the array nums, for each nums[i] find out how many numbers in the array are smaller than it. 
That is, for each nums[i] you have to count the number of valid j's such that j != i and nums[j] < nums[i].

Return the answer in an array.
*/


func smallerNumbersThanCurrent(nums []int) []int {
	res, buckets := make([]int, len(nums)), make([]int, 101)
	
	for _, num := range nums {
			buckets[num]++
	}
	
	sum := 0
	
	for i, bucket := range buckets {
			buckets[i], sum = sum, sum + bucket
	}
	
	for k, v := range nums {
			res[k] = buckets[v]
	}
	
	return res
}
