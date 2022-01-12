package intersection_of_two_arrays

func numberOfElementsInIntersection(arr1 []int, arr2 []int) int {
	occurrence := make(map[int]bool)
	for _, value := range arr1 {
		occurrence[value] = true
	}
	intersectedElementCount := 0
	for _, value := range arr2 {
		if occurrence[value] {
			intersectedElementCount = intersectedElementCount + 1
			delete(occurrence, value)
		}
	}
	return intersectedElementCount
}
