package trie

type Trie struct {
	nodeByCharacter map[rune]*Trie
	endOfWord       bool
}

func Constructor() Trie {
	return Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: false}
}

func (trie *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	head := trie
	for _, char := range word {
		if node, ok := head.nodeByCharacter[char]; ok {
			head = node
		} else {
			head.nodeByCharacter[char] = &Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: false}
			head = head.nodeByCharacter[char]
		}
	}
	head.endOfWord = true
}

func (trie *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	head := trie
	for _, char := range word {
		if node, ok := head.nodeByCharacter[char]; ok {
			head = node
		} else {
			return false
		}
	}
	return head.endOfWord == true
}

func (trie *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return false
	}
	head := trie
	for _, char := range prefix {
		if node, ok := head.nodeByCharacter[char]; ok {
			head = node
		} else {
			return false
		}
	}
	return true
}
