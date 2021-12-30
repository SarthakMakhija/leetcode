package word_dictionary

type TrieNode struct {
	nodeByCharacter map[rune]*TrieNode
	endOfWord       bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{nodeByCharacter: map[rune]*TrieNode{}, endOfWord: false},
	}
}

func (wordDictionary *WordDictionary) AddWord(word string) {
	head := wordDictionary.root
	for _, char := range word {
		node, ok := head.nodeByCharacter[char]
		if ok {
			head = node
		} else {
			head.nodeByCharacter[char] = &TrieNode{endOfWord: false, nodeByCharacter: map[rune]*TrieNode{}}
			head = head.nodeByCharacter[char]
		}
	}
	head.endOfWord = true
}

func (wordDictionary *WordDictionary) Search(word string) bool {
	const dot rune = 46

	var searchInner func(node *TrieNode, word string) bool
	searchInner = func(node *TrieNode, word string) bool {
		if len(word) == 0 {
			return node.endOfWord == true
		}
		char := rune(word[0])
		if char != dot {
			targetNode, ok := node.nodeByCharacter[char]
			if !ok {
				return false
			}
			return searchInner(targetNode, word[1:])
		} else {
			for _, targetNode := range node.nodeByCharacter {
				found := searchInner(targetNode, word[1:])
				if found {
					return true
				}
			}
			return false
		}
	}
	return searchInner(wordDictionary.root, word)
}
