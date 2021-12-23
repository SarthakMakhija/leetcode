package unique_paths

type position struct {
	row, col int
}

func uniquePaths(m int, n int) int {
	var uniquePathCache = make(map[position]int)
	var uniquePathsInner func(currentRow, currentCol int) int

	uniquePathsInner = func(currentRow, currentCol int) int {
		pathCount, ok := uniquePathCache[position{row: currentRow, col: currentCol}]
		switch {
		case ok:
			return pathCount
		case currentRow == m-1 && currentCol == n-1:
			return 1
		case currentRow >= m || currentCol >= n:
			return 0
		}
		uniquePathCache[position{row: currentRow, col: currentCol}] =
			uniquePathsInner(currentRow+1, currentCol) + uniquePathsInner(currentRow, currentCol+1)
		return uniquePathCache[position{row: currentRow, col: currentCol}]
	}
	return uniquePathsInner(0, 0)
}
