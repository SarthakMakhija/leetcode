package word_search

type position struct {
	row, col int
}

func exist(board [][]byte, original string) bool {
	rows := len(board)
	cols := len(board[0])

	var existInner func(row, col int, remainingWord string, positions []position) bool
	existInner = func(row, col int, remainingWord string, positions []position) bool {
		positions = append(positions, position{row: row, col: col})
		switch {
		case len(remainingWord) == 0:
			return true
		case row >= rows || col >= cols:
			return false
		case row < 0 || col < 0:
			return false
		case board[row][col] != remainingWord[0]:
			return false
		case len(remainingWord) == 1 && board[row][col] == remainingWord[0]:
			return true
		}
		nextPositions := neighborsMatchingCharacterAtIndex(board, remainingWord, 1, position{row: row, col: col})
		for _, nextPosition := range nextPositions {
			if !isVisited(nextPosition, positions) {
				if existInner(nextPosition.row, nextPosition.col, remainingWord[1:], positions) {
					return true
				}
			}
		}
		return false
	}
	for _, currentPosition := range firstCharacterIndicesIn(board, original) {
		if existInner(currentPosition.row, currentPosition.col, original, []position{}) {
			return true
		}
	}
	return false
}

func firstCharacterIndicesIn(board [][]byte, word string) []position {
	var indices []position
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == word[0] {
				indices = append(indices, position{row: row, col: col})
			}
		}
	}
	return indices
}

func neighborsMatchingCharacterAtIndex(board [][]byte, word string, index int, currentPosition position) []position {
	neighborPositions := []position{
		{row: currentPosition.row + 1, col: currentPosition.col},
		{row: currentPosition.row - 1, col: currentPosition.col},
		{row: currentPosition.row, col: currentPosition.col + 1},
		{row: currentPosition.row, col: currentPosition.col - 1},
	}

	var matchingPositions []position
	if index < len(word) {
		for _, neighborPosition := range neighborPositions {
			if isPositionInBounds(board, neighborPosition) && board[neighborPosition.row][neighborPosition.col] == word[index] {
				matchingPositions = append(matchingPositions, neighborPosition)
			}
		}
	}
	return matchingPositions
}

func isPositionInBounds(board [][]byte, currentPosition position) bool {
	rows, cols := len(board), len(board[0])
	switch {
	case currentPosition.row < 0 || currentPosition.col < 0:
		return false
	case currentPosition.row >= rows || currentPosition.col >= cols:
		return false
	}
	return true
}

func isVisited(pos position, positions []position) bool {
	for _, existingPosition := range positions {
		if existingPosition.row == pos.row && existingPosition.col == pos.col {
			return true
		}
	}
	return false
}
