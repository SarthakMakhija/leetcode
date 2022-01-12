package check_two_arrays_are_equal

func check(A []int, B []int) bool {
	if len(A) != len(B) {
		return false
	}
	common := make(map[int]int)
	for _, value := range A {
		common[value] = common[value] + 1
	}
	for _, value := range B {
		if _, ok := common[value]; !ok {
			return false
		} else {
			common[value] = common[value] - 1
		}
		if common[value] == 0 {
			delete(common, value)
		}
	}
	return len(common) == 0
}
