package zigzag_conversion

func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	firstRowNextIndex := numRows + (numRows - 1) - 1
	firstRowIndexDifference := firstRowNextIndex - 0
	nextCharacterIndex := firstRowIndexDifference

	result := ""
	for row := 1; row <= numRows; row++ {
		if row == 1 || row == numRows {
			result = result + convertForFirstAndLastRow(row, firstRowIndexDifference, s)
		} else {
			result = result + convertForRow(row, firstRowIndexDifference, nextCharacterIndex, s)
		}
		nextCharacterIndex = nextCharacterIndex - 1
	}
	return result
}

func convertForFirstAndLastRow(row int, indexDifference int, str string) string {
	output := ""
	characterIndex := row - 1

	for characterIndex < len(str) {
		output = output + string(str[characterIndex])
		characterIndex = characterIndex + indexDifference
	}
	return output
}

func convertForRow(row int, firstRowIndexDifference int, nextCharacterIndex int, str string) string {
	characterIndex := row - 1
	rowIndexDifference := nextCharacterIndex - characterIndex
	output := string(str[characterIndex])

	characterIndex = rowIndexDifference + characterIndex
	rowIndexDifference = firstRowIndexDifference - rowIndexDifference

	for characterIndex < len(str) {
		output = output + string(str[characterIndex])

		characterIndex = rowIndexDifference + characterIndex
		rowIndexDifference = firstRowIndexDifference - rowIndexDifference
	}
	return output
}
