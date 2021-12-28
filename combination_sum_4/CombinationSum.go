package combination_sum_4

func combinationSum4(candidates []int, target int) int {
	var combinationSumInner func(remaining int) int

	combinationSumCache := make(map[int]int)
	combinationSumInner = func(remaining int) int {
		cachedCombination, ok := combinationSumCache[remaining]
		switch {
		case ok:
			return cachedCombination
		case remaining == 0:
			return 1
		case remaining < 0:
			return -1
		}

		combinationCount := 0
		for _, value := range candidates {
			localCombinationCount := combinationSumInner(remaining - value)
			if localCombinationCount != -1 {
				combinationCount = combinationCount + localCombinationCount
			}
		}
		combinationSumCache[remaining] = combinationCount
		return combinationSumCache[remaining]
	}
	return combinationSumInner(target)
}
