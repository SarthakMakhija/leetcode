package decode_string

import "strconv"

func decodeString(s string) string {
	openingSquareBracket, closingSquareBracket := int32(91), int32(93)
	zeroUnicodePoint, nineUnicodePoint := int32(48), int32(57)

	numberAsString, countStack, resultStack := "", newStack(), newStack()
	for _, ch := range s {
		switch {
		case ch == openingSquareBracket:
			count, _ := strconv.Atoi(numberAsString)
			numberAsString = ""
			countStack.push(count)
			resultStack.push(string(ch))
		case ch == closingSquareBracket:
			count := countStack.pop().(int)
			part := resultStack.popTill(func(part string) bool {
				return part != "["
			})
			resultStack.push(repeat(part, count))
		case ch >= zeroUnicodePoint && ch <= nineUnicodePoint:
			numberAsString = numberAsString + string(ch)
		default:
			resultStack.push(string(ch))
		}
	}
	return resultStack.combineAsString()
}

func repeat(part string, count int) string {
	parts := ""
	for iteration := 1; iteration <= count; iteration++ {
		parts = parts + part
	}
	return parts
}

type Stack struct {
	elements []interface{}
}

func newStack() *Stack {
	return &Stack{elements: []interface{}{}}
}

func (stack *Stack) push(element interface{}) {
	stack.elements = append(stack.elements, element)
}

func (stack *Stack) popTill(f func(string) bool) string {
	parts := ""
	part := stack.pop().(string)
	for f(part) {
		parts = part + parts
		part = stack.pop().(string)
	}
	return parts
}

func (stack *Stack) pop() interface{} {
	element := stack.get()
	if len(stack.elements) == 0 {
		return element
	}
	stack.elements = stack.elements[0 : len(stack.elements)-1]
	return element
}

func (stack *Stack) get() interface{} {
	length := len(stack.elements)
	if length == 0 {
		return 0
	}
	element := stack.elements[length-1]
	return element
}

func (stack *Stack) combineAsString() string {
	parts, length := "", len(stack.elements)

	for count := 1; count <= length; count++ {
		part := stack.pop().(string)
		parts = part + parts
	}
	return parts
}
