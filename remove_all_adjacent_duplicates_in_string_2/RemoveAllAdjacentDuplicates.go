package remove_all_adjacent_duplicates_in_string_2

func removeDuplicates(str string, k int) string {
	if len(str) == 0 || k == 0 {
		return str
	}

	stack := newRuneStack()
	for _, ch := range str {
		if ch != stack.getAtTop() {
			stack.push(ch)
		} else {
			stack.pushAdjacentDuplicate(ch)
			if stack.top()-stack.beginIndexOf(ch)+1 == k {
				stack.popKTimes(k)
			}
		}
	}
	return stack.string()
}

type runeStack struct {
	elements              []rune
	characterBeginIndexes map[rune][]int
}

func newRuneStack() *runeStack {
	return &runeStack{elements: []rune{}, characterBeginIndexes: map[rune][]int{}}
}

func (stack *runeStack) push(element rune) {
	stack.elements = append(stack.elements, element)
	stack.characterBeginIndexes[element] = append(stack.characterBeginIndexes[element], len(stack.elements)-1)
}

func (stack *runeStack) pushAdjacentDuplicate(element rune) {
	stack.elements = append(stack.elements, element)
}

func (stack *runeStack) pop() rune {
	element := stack.getAtTop()

	if len(stack.elements) == 0 {
		return element
	}
	stack.elements = stack.elements[0 : len(stack.elements)-1]
	return element
}

func (stack *runeStack) popKTimes(k int) rune {
	element := stack.getAtTop()
	for count := 1; count <= k; count++ {
		stack.pop()
	}
	indexes := stack.characterBeginIndexes[element]
	if len(indexes) != 0 {
		stack.characterBeginIndexes[element] = indexes[0 : len(indexes)-1]
	}
	return element
}

func (stack *runeStack) getAtTop() rune {
	length := len(stack.elements)
	if length == 0 {
		return -1
	}
	element := stack.elements[length-1]
	return element
}

func (stack *runeStack) top() int {
	return len(stack.elements) - 1
}

func (stack *runeStack) beginIndexOf(ch rune) int {
	indexes := stack.characterBeginIndexes[ch]
	return indexes[len(indexes)-1]
}

func (stack *runeStack) string() string {
	return string(stack.elements)
}

func (stack *runeStack) isEmpty() bool {
	return len(stack.elements) == 0
}
