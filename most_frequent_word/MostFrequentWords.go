package most_frequent_word

func mostFrequentWord(words []string) string {
	trie := NewTrie()
	return trie.mostFrequentWord(words)
}

type Trie struct {
	nodeByCharacter      map[rune]*Trie
	endOfWord            bool
	occurrenceCount      int
	firstOccurrenceIndex int
}

func NewTrie() Trie {
	return Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: false}
}

func (trie *Trie) mostFrequentWord(words []string) string {
	mostFrequentOccurringWord, maxOccurrenceCount, insertIndex := "", -1, -1
	for index, word := range words {
		occurrenceCount, firstOccurrenceIndex := trie.insert(word, index)
		if (occurrenceCount > maxOccurrenceCount) ||
			(occurrenceCount == maxOccurrenceCount && firstOccurrenceIndex > insertIndex) {
			mostFrequentOccurringWord, maxOccurrenceCount, insertIndex = word, occurrenceCount, firstOccurrenceIndex
		}
	}
	return mostFrequentOccurringWord
}

func (trie *Trie) insert(word string, index int) (int, int) {
	if len(word) == 0 {
		return 0, 0
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
		head.firstOccurrenceIndex = index
	} else if !newInsert {
		head.occurrenceCount = head.occurrenceCount + 1
	}
	return head.occurrenceCount, head.firstOccurrenceIndex
}
