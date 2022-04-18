package rat_in_a_maze

type position struct {
	row, col  int
	direction string
}

type possibleDirections struct {
	d        []string
	possible bool
}

func findPath(maze [][]int, boardSize int) []string {
	if boardSize == 0 {
		return []string{}
	}

	isInBounds := func(pos position) bool {
		if pos.row < boardSize && pos.col < boardSize && pos.row >= 0 && pos.col >= 0 {
			return true
		}
		return false
	}
	isPositionUnblocked := func(pos position) bool {
		return maze[pos.row][pos.col] == 1
	}
	isAlreadyVisited := func(pos position, visitedPositions []position) bool {
		for _, visited := range visitedPositions {
			if pos.row == visited.row && pos.col == visited.col {
				return true
			}
		}
		return false
	}
	addDirection := func(directions []string, direction string) []string {
		if len(directions) == 0 {
			return []string{direction}
		}
		var newDirections []string
		for _, d := range directions {
			newDirections = append(newDirections, direction+d)
		}
		return newDirections
	}

	var pathFinderCache = make(map[position]possibleDirections)

	var findPathInner func(currentPosition position, visitedPositions []position) possibleDirections
	findPathInner = func(currentPosition position, visitedPositions []position) possibleDirections {
		visitedPositions = append(visitedPositions, currentPosition)
		cached, ok := pathFinderCache[currentPosition]
		switch {
		case ok:
			return cached
		case currentPosition.row == boardSize-1 && currentPosition.col == boardSize-1:
			return possibleDirections{d: []string{}, possible: true}
		}
		nextPositions := []position{
			{row: currentPosition.row + 1, col: currentPosition.col, direction: "D"},
			{row: currentPosition.row - 1, col: currentPosition.col, direction: "U"},
			{row: currentPosition.row, col: currentPosition.col + 1, direction: "R"},
			{row: currentPosition.row, col: currentPosition.col - 1, direction: "L"},
		}
		var directions []string
		for _, nextPosition := range nextPositions {
			if isInBounds(nextPosition) &&
				!isAlreadyVisited(nextPosition, visitedPositions) &&
				isPositionUnblocked(nextPosition) {
				possibleDirections := findPathInner(nextPosition, visitedPositions)
				if possibleDirections.possible {
					directions = append(directions, addDirection(possibleDirections.d, nextPosition.direction)...)
				}
			}
		}
		switch {
		case len(directions) == 0:
			pd := possibleDirections{[]string{}, false}
			pathFinderCache[currentPosition] = pd
			return pd
		default:
			pd := possibleDirections{directions, true}
			pathFinderCache[currentPosition] = pd
			return pd
		}
	}
	return findPathInner(position{row: 0, col: 0}, []position{}).d
}
