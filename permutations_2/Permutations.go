package permutations_2

func permuteUnique(nums []int) [][]int {
	var permutationsInner func(numbers []int) [][]int
	permutationsInner = func(numbers []int) [][]int {
		if len(numbers) == 1 {
			return [][]int{numbers}
		} else if len(numbers) == 2 {
			if numbers[0] != numbers[1] {
				return [][]int{numbers, {numbers[1], numbers[0]}}
			}
			return [][]int{numbers}
		}
		var allPermutations [][]int
		if len(numbers) > 2 {
			for index := 0; index < len(numbers); index++ {
				if index != 0 {
					numbers[0], numbers[index] = numbers[index], numbers[0]
				}
				allPermutations = merge(allPermutations, prepend(permutationsInner(numbers[1:]), numbers[0]))
				if index != 0 {
					numbers[index], numbers[0] = numbers[0], numbers[index]
				}
			}
		}
		return allPermutations
	}
	return permutationsInner(nums)
}

func prepend(permutations [][]int, element int) [][]int {
	var newPermutations [][]int
	for index := 0; index < len(permutations); index++ {
		newPermutations = append(newPermutations, append([]int{element}, permutations[index]...))
	}
	return newPermutations
}

func merge(one [][]int, other [][]int) [][]int {
	var merged [][]int
	for index := 0; index < len(one); index++ {
		if !contains(merged, one[index]) {
			merged = append(merged, one[index])
		}
	}
	for index := 0; index < len(other); index++ {
		if !contains(merged, other[index]) {
			merged = append(merged, other[index])
		}
	}
	return merged
}

func contains(merged [][]int, elements []int) bool {
	lengthElements := len(elements)
	for index := 0; index < len(merged); index++ {
		if len(merged[index]) == lengthElements {
			iterator := 0
			for ; iterator < lengthElements; iterator++ {
				if merged[index][iterator] != elements[iterator] {
					break
				}
			}
			if iterator >= lengthElements {
				return true
			}
		}
	}
	return false
}
