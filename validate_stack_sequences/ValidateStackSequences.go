package validate_stack_sequences

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 || len(popped) == 0 {
		return false
	}

	stack, popIndex := newIntStack(), 0
	for _, pushedElement := range pushed {
		stack.push(pushedElement)
		for !stack.isEmpty() && stack.get() == popped[popIndex] {
			stack.pop()
			popIndex = popIndex + 1
		}
	}
	if stack.isEmpty() {
		return true
	}
	return false
}

type intStack struct {
	elements []int
}

func newIntStack() *intStack {
	return &intStack{elements: []int{}}
}

func (stack *intStack) push(element int) {
	stack.elements = append(stack.elements, element)
}

func (stack *intStack) pop() int {
	element := stack.get()
	if len(stack.elements) == 0 {
		return element
	}
	stack.elements = stack.elements[0 : len(stack.elements)-1]
	return element
}

func (stack *intStack) get() int {
	length := len(stack.elements)
	if length == 0 {
		return 0
	}
	element := stack.elements[length-1]
	return element
}

func (stack *intStack) isEmpty() bool {
	return len(stack.elements) == 0
}
