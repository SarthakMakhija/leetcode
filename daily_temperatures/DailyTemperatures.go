package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	}
	length := len(temperatures)
	temperatureIndexStack, daysToWait := newIntStack(), make([]int, length)

	temperatureIndexStack.push(length - 1)
	daysToWait[length-1] = 0
	for index := length - 2; index >= 0; index-- {
		for !temperatureIndexStack.isEmpty() && temperatures[temperatureIndexStack.get()] <= temperatures[index] {
			temperatureIndexStack.pop()
		}
		if temperatureIndexStack.isEmpty() {
			daysToWait[index] = 0
		} else {
			daysToWait[index] = temperatureIndexStack.get() - index
		}
		temperatureIndexStack.push(index)
	}
	return daysToWait
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
