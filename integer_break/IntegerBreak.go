package integer_break

func integerBreak(n int) int {
	if n <= 1 {
		return n
	}
	var integerBreakInner func(remaining, partCount int) (int, int)
	integerBreakCache := make(map[int]IntegerBreak)

	integerBreakInner = func(remaining, partCount int) (int, int) {
		cached, ok := integerBreakCache[remaining]
		switch {
		case ok:
			return cached.product, cached.partCount
		case remaining == 0:
			return 1, partCount
		case remaining < 0:
			return 0, 0
		}
		maxProduct, totalPartCount := -1, 0
		for value := 1; value <= remaining; value++ {
			localProduct, localPartCount := integerBreakInner(remaining-value, partCount+1)
			localProduct = localProduct * value

			if localProduct > maxProduct && localPartCount >= 2 {
				maxProduct, totalPartCount = localProduct, localPartCount
			}
		}
		integerBreakCache[remaining] = IntegerBreak{product: maxProduct, partCount: totalPartCount}
		return maxProduct, totalPartCount
	}
	product, _ := integerBreakInner(n, 0)
	return product
}

type IntegerBreak struct {
	product, partCount int
}
