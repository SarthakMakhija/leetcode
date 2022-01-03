package word_search

type position struct {
	row, col int
}

func exist(board [][]byte, word string) bool {
	var visitedPositions map[position]bool
	rows, cols := len(board), len(board[0])

	var existInner func(currentPosition position, word string) bool
	existInner = func(currentPosition position, word string) bool {
		visitedPositions[currentPosition] = true
		switch {
		case len(word) == 0:
			return true
		case len(word) == 1 && board[currentPosition.row][currentPosition.col] == word[0]:
			return true
		case currentPosition.row < 0 || currentPosition.col < 0:
			return false
		case currentPosition.row >= rows || currentPosition.col >= cols:
			return false
		case board[currentPosition.row][currentPosition.col] != word[0]:
			return false
		}
		nextPositions := neighborsMatchingCharacterAtIndex(board, word, 1, currentPosition)
		for _, nextPosition := range nextPositions {
			if !visitedPositions[nextPosition] {
				if existInner(nextPosition, word[1:]) {
					return true
				}
			}
		}
		delete(visitedPositions, currentPosition)
		return false
	}
	for _, currentPosition := range firstCharacterIndicesIn(board, word) {
		visitedPositions = make(map[position]bool)
		if existInner(currentPosition, word) {
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

	if index < len(word) {
		var matchingPositions []position
		for _, neighborPosition := range neighborPositions {
			if isPositionInBounds(board, neighborPosition) && board[neighborPosition.row][neighborPosition.col] == word[index] {
				matchingPositions = append(matchingPositions, neighborPosition)
			}
		}
		return matchingPositions
	}
	return []position{}
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
