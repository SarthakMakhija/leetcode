package next_greater_node_in_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var nextGreaterNodeInner func(node *ListNode)
	var greaterNodeValues []int
	stack := newIntStack()

	nextGreaterNodeInner = func(node *ListNode) {
		if node.Next == nil {
			stack.push(node.Val)
			greaterNodeValues = append(greaterNodeValues, 0)
			return
		}
		nextGreaterNodeInner(node.Next)
		for !stack.isEmpty() && stack.get() <= node.Val {
			stack.pop()
		}
		if stack.isEmpty() {
			greaterNodeValues = append(greaterNodeValues, 0)
		} else {
			greaterNodeValues = append(greaterNodeValues, stack.get())
		}
		stack.push(node.Val)
	}
	nextGreaterNodeInner(head)
	return reverse(greaterNodeValues)
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
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
