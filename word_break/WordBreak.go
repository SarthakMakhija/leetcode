package word_break

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	var wordBreakInner func(remaining string) bool
	wordBreakCache := make(map[string]bool)

	wordBreakInner = func(remaining string) bool {
		canBeConstructed, ok := wordBreakCache[remaining]
		switch {
		case ok:
			return canBeConstructed
		case len(remaining) == 0:
			return true
		}
		for _, word := range wordDict {
			if strings.Index(remaining, word) == 0 {
				if wordBreakInner(remaining[len(word):]) {
					wordBreakCache[remaining] = true
					return true
				}
			}
		}
		wordBreakCache[remaining] = false
		return false
	}
	return wordBreakInner(s)
}
