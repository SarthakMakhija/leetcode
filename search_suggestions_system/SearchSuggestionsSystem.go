package search_suggestions_system

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := newTrie()
	addProducts := func() {
		for _, product := range products {
			trie.add(product)
		}
	}

	var allMatchingWords [][]string
	searchProducts := func() {
		for index := range searchWord {
			matchingWords := trie.searchWordsStartingWith(searchWord[0 : index+1])
			allMatchingWords = append(allMatchingWords, matchingWords)
		}
	}

	addProducts()
	searchProducts()
	return allMatchingWords
}

type trie struct {
	root *trieNode
}

type trieNode struct {
	characters  [26]*trieNode
	isEndOfWord bool
}

func newTrie() *trie {
	return &trie{root: &trieNode{}}
}

func (t *trie) add(word string) {
	head := t.root
	for _, ch := range word {
		characterIndex := ch - 97
		if head.characters[characterIndex] == nil {
			head.characters[characterIndex] = &trieNode{isEndOfWord: false}
		}
		head = head.characters[characterIndex]
	}
	head.isEndOfWord = true
}

func (t *trie) searchWordsStartingWith(prefix string) []string {
	head := t.root
	for _, ch := range prefix {
		characterIndex := ch - 97
		if head.characters[characterIndex] == nil {
			return []string{}
		}
		head = head.characters[characterIndex]
	}

	var matchingWords []string
	var searchInner func(node *trieNode, word string)

	searchInner = func(node *trieNode, word string) {
		if node.isEndOfWord {
			matchingWords = append(matchingWords, word)
		}
		for index, next := range node.characters {
			if next != nil {
				if len(matchingWords) < 3 {
					searchInner(next, word+string(rune(97+index)))
				} else {
					break
				}
			}
		}
	}
	searchInner(head, prefix)
	return matchingWords
}
