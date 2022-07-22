/*
You are given four integers row, cols, rCenter, and cCenter. 
There is a rows x cols matrix and you are on the cell with the coordinates (rCenter, cCenter).

Return the coordinates of all cells in the matrix, sorted by their distance from (rCenter, cCenter) from the smallest distance to the largest distance. 
You may return the answer in any order that satisfies this condition.
The distance between two cells (r1, c1) and (r2, c2) is |r1 - r2| + |c1 - c2|.
*/

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	buckets, res := [][][]int{}, [][]int{}
	
	for i := 0; i < (rows + cols - 1); i++ {
			buckets = append(buckets, [][]int{})
	}
	
	for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
					distance := abs(r - rCenter) + abs(c - cCenter)
					buckets[distance] = append(buckets[distance], []int{r, c})
			}
	}
	
	for _, bucket := range buckets {
			res = append(res, bucket...)
	}
	
	return res
}

func abs(x int) int {
	if x < 0 {
			return x * -1
	}
	return x
}
