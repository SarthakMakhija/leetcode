package coin_change

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	var coinChangeInner func(remaining int) int

	coinChangeCache := make(map[int]int)
	coinChangeInner = func(remaining int) int {
		cachedCoins, ok := coinChangeCache[remaining]
		switch {
		case ok:
			return cachedCoins
		case remaining < 0:
			return math.MaxInt32 - 1
		case remaining == 0:
			return 0
		}
		minimumCoins := -1
		for _, coin := range coins {
			totalCoins := coinChangeInner(remaining - coin)
			if totalCoins != -1 {
				totalCoins = totalCoins + 1
			}
			if totalCoins != math.MaxInt32 && totalCoins != -1 {
				if minimumCoins == -1 || totalCoins < minimumCoins {
					minimumCoins = totalCoins
				}
			}
		}
		coinChangeCache[remaining] = minimumCoins
		return minimumCoins
	}
	return coinChangeInner(amount)
}
