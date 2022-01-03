package integer_replacement

func integerReplacement(n int) int {
	var integerReplacementInner func(remaining int) int
	integerReplacementCache := make(map[int]int)

	integerReplacementInner = func(remaining int) int {
		cached, ok := integerReplacementCache[remaining]
		switch {
		case ok:
			return cached
		case remaining == 1:
			return 0
		}
		if remaining%2 == 0 {
			integerReplacementCache[remaining] = integerReplacementInner(remaining/2) + 1
			return integerReplacementCache[remaining]
		} else {
			integerReplacementCache[remaining] = minOf(integerReplacementInner(remaining+1)+1, integerReplacementInner(remaining-1)+1)
			return integerReplacementCache[remaining]
		}
	}
	return integerReplacementInner(n)
}

func minOf(a, b int) int {
	if a < b {
		return a
	}
	return b
}
