package word_break

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}

	var wordBreakInner func(remaining string) bool
	var wordBreakCache = make(map[string]bool)

	wordBreakInner = func(remaining string) bool {
		breakable, ok := wordBreakCache[remaining]
		switch {
		case len(remaining) == 0:
			return true
		case ok:
			return breakable
		}
		isBreakPossible := false
		for _, word := range wordDict {
			if strings.HasPrefix(remaining, word) {
				isBreakPossible = wordBreakInner(remaining[len(word):])
				if isBreakPossible {
					break
				}
			}
		}
		wordBreakCache[remaining] = isBreakPossible
		return wordBreakCache[remaining]
	}
	return wordBreakInner(s)
}
