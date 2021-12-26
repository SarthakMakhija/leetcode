package triangle

import "math"

type position struct {
	row, col int
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	var minimumTotalInner func(row int, index int) int

	totalCache := make(map[position]int)
	rows := len(triangle)
	minimumTotalInner = func(row int, index int) int {
		total, ok := totalCache[position{row: row, col: index}]
		switch {
		case ok:
			return total
		case row == rows-1 && index < len(triangle[row]):
			return triangle[row][index]
		case row >= rows:
			return math.MaxInt32
		}
		totalCache[position{row: row, col: index}] = minimum(
			triangle[row][index]+minimumTotalInner(row+1, index),
			triangle[row][index]+minimumTotalInner(row+1, index+1),
		)
		return totalCache[position{row: row, col: index}]
	}
	return minimumTotalInner(0, 0)
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
