package rpn

import "strconv"

type IntStack struct {
	elements []int
}

func newIntStack() *IntStack {
	return &IntStack{elements: []int{}}
}

func (stack *IntStack) push(element int) {
	stack.elements = append(stack.elements, element)
}

func (stack *IntStack) pop() int {
	element := stack.get()
	if len(stack.elements) == 0 {
		return element
	}
	stack.elements = stack.elements[0 : len(stack.elements)-1]
	return element
}

func (stack *IntStack) get() int {
	length := len(stack.elements)
	if length == 0 {
		return 0
	}
	element := stack.elements[length-1]
	return element
}

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return -1
	}
	operationByOperator := map[string]func(int, int) int{
		"+": func(addend int, augend int) int {
			return addend + augend
		},
		"/": func(numerator int, denominator int) int {
			return numerator / denominator
		},
		"-": func(a int, b int) int {
			return a - b
		},
		"*": func(a int, b int) int {
			return a * b
		},
	}
	stack := newIntStack()
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			first := stack.pop()
			second := stack.pop()
			stack.push(operationByOperator[token](second, first))
		} else {
			intVal, _ := strconv.Atoi(token)
			stack.push(intVal)
		}
	}
	return stack.get()
}
