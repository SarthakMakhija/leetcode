package search_2D_matrix

import (
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	rowFound, rowToSearchIndex, rows := false, -1, len(matrix)
	for row := 0; row < rows; row++ {
		if target < matrix[row][0] {
			rowFound, rowToSearchIndex = true, row-1
			break
		}
	}
	if !rowFound {
		rowToSearchIndex = rows - 1
	}
	if rowToSearchIndex < 0 {
		return false
	}
	rowToSearch := matrix[rowToSearchIndex]
	targetIndex := sort.Search(len(rowToSearch), func(index int) bool {
		return rowToSearch[index] >= target
	})
	if targetIndex < len(rowToSearch) && rowToSearch[targetIndex] == target {
		return true
	}
	return false
}
