package first_repeating_element

func firstRepeated(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	lowestIndex, frequency := -1, make(map[int]CountByStartingIndex)
	for index, value := range arr {
		if countByStartingIndex, ok := frequency[value]; ok {
			frequency[value] = CountByStartingIndex{count: countByStartingIndex.count + 1, index: countByStartingIndex.index}
			updatedCountByStartingIndex := frequency[value]
			if updatedCountByStartingIndex.count >= 2 {
				if lowestIndex == -1 || updatedCountByStartingIndex.index < lowestIndex {
					lowestIndex = updatedCountByStartingIndex.count
				}
			}
		} else {
			frequency[value] = CountByStartingIndex{count: 1, index: index + 1}
		}
	}
	return lowestIndex
}

type CountByStartingIndex struct {
	count, index int
}
