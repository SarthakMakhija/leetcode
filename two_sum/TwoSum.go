package two_sum

func twoSum(nums []int, target int) []int {

	indexByElement := make(map[int]int)
	for index, element := range nums {
		if existingIndex, ok := indexByElement[target-element]; ok {
			return []int{existingIndex, index}
		}
		indexByElement[element] = index
	}
	return []int{}
}
