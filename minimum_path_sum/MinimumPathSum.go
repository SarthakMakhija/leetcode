package minimum_path_sum

import "math"

type position struct {
	row, col int
}

func minPathSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var uniquePathCache = make(map[position]int)
	var minimumPathSumInner func(row, col int) int

	minimumPathSumInner = func(row, col int) int {
		minimumCostPath, ok := uniquePathCache[position{row: row, col: col}]
		switch {
		case ok:
			return minimumCostPath
		case row == rows-1 && col == cols-1:
			return grid[row][col]
		case row >= rows || col >= cols:
			return math.MaxInt32
		}
		uniquePathCache[position{row: row, col: col}] = min(
			grid[row][col]+minimumPathSumInner(row+1, col),
			grid[row][col]+minimumPathSumInner(row, col+1),
		)
		return uniquePathCache[position{row: row, col: col}]
	}
	return minimumPathSumInner(0, 0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
