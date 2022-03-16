package next_greater_element_2

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	originalLength := len(nums)
	circular := make([]int, originalLength+originalLength-1)
	copy(circular, nums)
	copy(circular[originalLength:], nums[0:originalLength-1])

	circularLength := len(circular)
	stack, nextGreater := newIntStack(), make([]int, circularLength)

	stack.push(circular[circularLength-1])
	nextGreater[len(nextGreater)-1] = -1

	for index := circularLength - 2; index >= 0; index-- {
		for !stack.isEmpty() && stack.get() <= circular[index] {
			stack.pop()
		}
		if stack.isEmpty() {
			nextGreater[index] = -1
		} else {
			nextGreater[index] = stack.get()
		}
		stack.push(circular[index])
	}
	return nextGreater[0:originalLength]
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
