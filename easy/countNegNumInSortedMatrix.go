/*
Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, 
return the number of negative numbers in grid.
*/

func countNegatives(grid [][]int) int {
	//loop inner arrays in reverse
	count := 0
	for _, val := range grid {
		//check first element in array - if yes add length of array to count
		mid := len(val) / 2
		if val[0] < 0 {
			count += len(val)
		} else if len(val) > mid+1 && val[mid+1] < 0 && val[mid] >= 0 {
			//check mod + 1 element
			// if -ve add length of subarray
			count += len(val[mid+1:])
		} else {
			for i := len(val) - 1; i >= 0; i-- {
				if val[i] < 0 {
					count += 1
				}
			}
		}
	}
	return count
}
