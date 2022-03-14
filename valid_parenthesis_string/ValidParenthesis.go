package valid_parenthesis_string

type characterPosition struct {
	character rune
	position  int
}

func checkValidString(str string) bool {
	if len(str) == 0 {
		return false
	}
	parenthesisStack, wildCardStack := newCharacterPositionStack(), newCharacterPositionStack()
	for index, ch := range str {
		switch ch {
		case '(':
			parenthesisStack.push(characterPosition{character: ch, position: index})
		case '*':
			wildCardStack.push(characterPosition{character: ch, position: index})
		case ')':
			topCharacterPosition := parenthesisStack.pop()
			if topCharacterPosition.position == -1 {
				topWildCardPosition := wildCardStack.pop()
				if topWildCardPosition.position == -1 {
					return false
				}
			}
		default:
			return false
		}
	}
	for !parenthesisStack.isEmpty() {
		parenthesis := parenthesisStack.pop()
		wildCard := wildCardStack.pop()
		if wildCard.position == -1 || wildCard.position < parenthesis.position {
			return false
		}
	}
	return true
}

type characterPositionStack struct {
	elements []characterPosition
}

func newCharacterPositionStack() *characterPositionStack {
	return &characterPositionStack{elements: []characterPosition{}}
}

func (stack *characterPositionStack) push(element characterPosition) {
	stack.elements = append(stack.elements, element)
}

func (stack *characterPositionStack) pop() characterPosition {
	element := stack.get()
	if len(stack.elements) == 0 {
		return element
	}
	stack.elements = stack.elements[0 : len(stack.elements)-1]
	return element
}

func (stack *characterPositionStack) get() characterPosition {
	length := len(stack.elements)
	if length == 0 {
		return characterPosition{position: -1}
	}
	element := stack.elements[length-1]
	return element
}

func (stack *characterPositionStack) all() []characterPosition {
	return stack.elements
}

func (stack *characterPositionStack) isEmpty() bool {
	return len(stack.elements) == 0
}
