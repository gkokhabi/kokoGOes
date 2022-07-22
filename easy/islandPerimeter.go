/*
You are given row x col grid representing a map where grid[i][j] = 1 represents land and grid[i][j] = 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). 
The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes", meaning the water inside isn't connected to the water around the island. 
One cell is a square with side length 1. 
The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.
*/

func islandPerimeter(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dirs := []int{1, 0, -1, 0, 1}

	ans := 0

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x > m-1 || y < 0 || y > n-1 || grid[x][y] == 0 {
			ans++
			return
		}
		if grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2
		for i := 0; i < 4; i++ {
			xi := x + dirs[i]
			yi := y + dirs[i+1]
			dfs(xi, yi)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
			}
		}
	}
	fmt.Println(grid)
	return ans
}
