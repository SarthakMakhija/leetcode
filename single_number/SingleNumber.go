package single_number

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result = result ^ num
	}
	return result
}
