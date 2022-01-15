package camel_case

func findAllWords(dict []string, pattern string) []string {
	trie := NewTrie()
	trie.buildDictionary(dict)
	return trie.findMatchingCamelCaseWords(pattern)
}

type Trie struct {
	nodeByCharacter map[rune]*Trie
	endOfWord       bool
	occurrenceCount int
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
	newInsert, head := false, trie
	for _, char := range word {
		if _, ok := head.nodeByCharacter[char]; !ok {
			head.nodeByCharacter[char] = &Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: false}
			newInsert = true
		}
		head = head.nodeByCharacter[char]
	}
	head.endOfWord = true
	if head.occurrenceCount == 0 {
		head.occurrenceCount = 1
	} else if !newInsert {
		head.occurrenceCount = head.occurrenceCount + 1
	}
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
			for count := 1; count <= node.occurrenceCount; count++ {
				words = append(words, word)
			}
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
