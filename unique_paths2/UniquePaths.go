package unique_paths2

type position struct {
	row, col int
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var uniquePathsWithObstaclesInner func(row, col int) int
	uniquePathsWithObstaclesCache := make(map[position]int)

	rows, cols := len(obstacleGrid), len(obstacleGrid[0])
	uniquePathsWithObstaclesInner = func(row, col int) int {
		pathCount, ok := uniquePathsWithObstaclesCache[position{row: row, col: col}]
		switch {
		case ok:
			return pathCount
		case row >= rows || col >= cols:
			return 0
		case row == rows-1 && col == cols-1 && obstacleGrid[row][col] == 0:
			return 1
		case obstacleGrid[row][col] == 1:
			return 0
		}
		uniquePathsWithObstaclesCache[position{row: row, col: col}] =
			uniquePathsWithObstaclesInner(row+1, col) + uniquePathsWithObstaclesInner(row, col+1)
		return uniquePathsWithObstaclesCache[position{row: row, col: col}]
	}
	return uniquePathsWithObstaclesInner(0, 0)
}
