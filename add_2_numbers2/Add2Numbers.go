package add_2_numbers2

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Stack, l2Stack := fillInStack(l1), fillInStack(l2)

	var newHead *ListNode = nil
	carry := 0

	for !l1Stack.isEmpty() || !l2Stack.isEmpty() {
		sum := l1Stack.pop() + l2Stack.pop() + carry
		carry = sum / 10
		node := &ListNode{Val: sum % 10}

		if newHead == nil {
			newHead = node
		} else {
			node.Next = newHead
			newHead = node
		}
	}
	if carry != 0 {
		node := &ListNode{Val: carry}
		node.Next = newHead
		newHead = node
	}
	return newHead
}

func (node *ListNode) all() []int {
	var values []int
	head := node

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}

func fillInStack(list *ListNode) *IntStack {
	head, stack := list, newIntStack()
	for head != nil {
		stack.push(head.Val)
		head = head.Next
	}
	return stack
}

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

func (stack *IntStack) isEmpty() bool {
	return len(stack.elements) == 0
}
