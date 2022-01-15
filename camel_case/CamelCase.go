package camel_case

func findAllWords(dict []string, pattern string) []string {
	trie := NewTrie()
	trie.buildDictionary(dict)
	return trie.findMatchingCamelCaseWords(pattern)
}

type Trie struct {
	nodeByCharacter map[rune]*Trie
	endOfWord       bool
}

func NewTrie() Trie {
	return Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: false}
}

func (trie *Trie) buildDictionary(words []string) {
	for _, word := range words {
		trie.Insert(word)
	}
}

func (trie *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	head := trie
	for _, char := range word {
		if _, ok := head.nodeByCharacter[char]; !ok {
			head.nodeByCharacter[char] = &Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: false}
		}
		head = head.nodeByCharacter[char]
	}
	head.endOfWord = true
}

func (trie *Trie) findMatchingCamelCaseWords(pattern string) []string {
	var words []string
	head, ok := trie.nodeByCharacter[rune(pattern[0])]
	if !ok {
		return words
	}
	var findMatchingCamelCaseWordsInner func(node *Trie, patternIndex int, word string)
	findMatchingCamelCaseWordsInner = func(node *Trie, patternIndex int, word string) {
		switch {
		case patternIndex >= len(pattern) && node.endOfWord:
			words = append(words, word)
		}
		for ch, targetNode := range node.nodeByCharacter {
			if patternIndex < len(pattern) && (ch >= 'A' && ch <= 'Z') && ch != rune(pattern[patternIndex]) {
				continue
			}
			nextIndex := patternIndex
			if patternIndex < len(pattern) && ch == rune(pattern[patternIndex]) {
				nextIndex = patternIndex + 1
			}
			findMatchingCamelCaseWordsInner(targetNode, nextIndex, word+string(ch))
		}
	}
	findMatchingCamelCaseWordsInner(head, 1, string(pattern[0]))
	return words
}
