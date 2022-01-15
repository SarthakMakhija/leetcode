package renaming_cities

import "strconv"

type uniqueCityCode = string

func rename(cities []string) []string {
	trie := NewTrie()
	var uniqueCodes []uniqueCityCode
	for _, city := range cities {
		uniqueCodes = append(uniqueCodes, trie.insert(city))
	}
	return uniqueCodes
}

//Trie
//We could have used occurrenceCount in this struct to simplify the insert method, but
//geeksforgeeks has already defined a TrieNode with just these 2 fields:
//Map<Character, TrieNode> & endOfWord
type Trie struct {
	nodeByCharacter map[rune]*Trie
	endOfWord       bool
}

func NewTrie() Trie {
	return Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: false}
}

func (trie *Trie) insert(word string) uniqueCityCode {
	if len(word) == 0 {
		return ""
	}
	cityCodeIndex, head := -1, trie
	for index, char := range word {
		if _, ok := head.nodeByCharacter[char]; !ok {
			head.nodeByCharacter[char] = &Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: false}
			if cityCodeIndex == -1 {
				cityCodeIndex = index
			}
		}
		head = head.nodeByCharacter[char]
	}
	if cityCodeIndex != -1 {
		head.endOfWord = true
		return word[0 : cityCodeIndex+1]
	} else if cityCodeIndex == -1 && !head.endOfWord {
		head.endOfWord = true
		return word
	}

	head.endOfWord = true
	nextOccurrenceCount := 2

	if _, ok := head.nodeByCharacter['#']; !ok {
		head.nodeByCharacter['#'] = &Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: false}
		head.nodeByCharacter['#'].nodeByCharacter[rune(nextOccurrenceCount)] = &Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: true}
	} else {
		nodeWithOccurrenceCount := head.nodeByCharacter['#']
		for occurrenceCount := range nodeWithOccurrenceCount.nodeByCharacter {
			delete(nodeWithOccurrenceCount.nodeByCharacter, occurrenceCount)
			nextOccurrenceCount = int(occurrenceCount) + 1
			break
		}
		nodeWithOccurrenceCount.nodeByCharacter[rune(nextOccurrenceCount)] = &Trie{nodeByCharacter: map[rune]*Trie{}, endOfWord: true}
	}
	return word + " " + strconv.Itoa(nextOccurrenceCount)
}
