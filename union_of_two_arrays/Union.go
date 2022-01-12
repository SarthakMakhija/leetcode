package union_of_two_arrays

func doUnion(arr1 []int, arr2 []int) int {
	union := make(map[int]bool)
	for _, value := range arr1 {
		union[value] = true
	}
	for _, value := range arr2 {
		union[value] = true
	}
	return len(union)
}
