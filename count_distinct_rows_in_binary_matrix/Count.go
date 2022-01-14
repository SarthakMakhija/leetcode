package count_distinct_rows_in_binary_matrix

func countDistinctRowsIn(matrix [][]bool) int {
	trie := newBooleanMatrixTrie()
	for _, row := range matrix {
		trie.insert(row)
	}
	return trie.count()
}

type BooleanMatrixTrie struct {
	nodeByBoolean map[bool]*BooleanMatrixTrie
	endOfRow      bool
}

func newBooleanMatrixTrie() *BooleanMatrixTrie {
	return &BooleanMatrixTrie{nodeByBoolean: map[bool]*BooleanMatrixTrie{}, endOfRow: false}
}

func (booleanMatrixTrie *BooleanMatrixTrie) insert(row []bool) {
	head := booleanMatrixTrie
	for _, value := range row {
		if _, ok := head.nodeByBoolean[value]; !ok {
			head.nodeByBoolean[value] =
				&BooleanMatrixTrie{nodeByBoolean: map[bool]*BooleanMatrixTrie{}, endOfRow: false}
		}
		head = head.nodeByBoolean[value]
	}
	head.endOfRow = true
}

func (booleanMatrixTrie *BooleanMatrixTrie) count() int {
	var countInner func(node *BooleanMatrixTrie) int
	countInner = func(node *BooleanMatrixTrie) int {
		if node.endOfRow {
			return 1
		}
		count := 0
		for _, targetNode := range node.nodeByBoolean {
			count = count + countInner(targetNode)
		}
		return count
	}
	return countInner(booleanMatrixTrie)
}
