package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	colIndexToSetZero := make(map[int]bool)
	for rowIndex := 0; rowIndex < len(matrix); rowIndex++ {
		foundZero := false
		for colIndex := 0; colIndex < len(matrix[rowIndex]); colIndex++ {
			if matrix[rowIndex][colIndex] == 0 {
				foundZero, colIndexToSetZero[colIndex] = true, true
				setPreviousColumnCellsToZero(rowIndex, colIndex, matrix)
				setPreviousRowCellsToZero(rowIndex, colIndex, matrix)
			}
			if foundZero || colIndexToSetZero[colIndex] {
				matrix[rowIndex][colIndex] = 0
			}
		}
	}
}

func setPreviousColumnCellsToZero(rowIndex, colIndex int, matrix [][]int) {
	for currentRowIndex := rowIndex - 1; currentRowIndex >= 0; currentRowIndex-- {
		matrix[currentRowIndex][colIndex] = 0
	}
}

func setPreviousRowCellsToZero(rowIndex, colIndex int, matrix [][]int) {
	for currentColumnIndex := colIndex - 1; currentColumnIndex >= 0; currentColumnIndex-- {
		matrix[rowIndex][currentColumnIndex] = 0
	}
}
