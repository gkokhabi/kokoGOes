/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers.
You may assume that each input would have exactly one solution.
*/

func valid(x int) int {
	if x < 0 return -x
	return x
}

func threeSumClosest(nums []int, target int) int {
    l := len(nums)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	// take arbitrary sum as closest
	closest := nums[0] + nums[1] + nums[2]

	for i := 0; i < l-2; i++ {
		// ideal sum for (2nd+3rd)
		c := target - nums[i]
		// start looking at ends
		l, h := i+1, l-1
		for l < h {
			v := nums[l] + nums[h]
			w := nums[i] + v
			// for each three numbers check if their sum is smaller than current closest number
			if valid(target-w) < valid(target-closest) {
				closest = w
			}
			if v > c {
				// if sum of (2nd+3rd) is larger that ideal,then decrease it
				h--
			} else if v < c {
				// if sum of (2nd+3rd) is smaller than ideal,then increase it
				l++
			} else {
				//otherwise exact match- then just return it
				return target
			}
		}
	}

	return closest
}


