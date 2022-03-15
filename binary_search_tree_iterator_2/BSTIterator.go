package binary_search_tree_iterator_2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	current   *TreeNode
	nodeStack *treeNodeStack
}

func Constructor(root *TreeNode) BSTIterator {
	stack, head := newTreeNodeStack(), root
	for head.Left != nil {
		stack.push(head)
		head = head.Left
	}
	iterator := BSTIterator{current: head, nodeStack: stack}
	return iterator
}

func (bstIterator *BSTIterator) Next() int {
	value := bstIterator.current.Val
	if bstIterator.current.Right == nil {
		bstIterator.current = bstIterator.nodeStack.pop()
	} else {
		head := bstIterator.current.Right
		for head.Left != nil {
			bstIterator.nodeStack.push(head)
			head = head.Left
		}
		bstIterator.current = head
	}
	return value
}

func (bstIterator *BSTIterator) HasNext() bool {
	return bstIterator.current != nil
}

type treeNodeStack struct {
	elements []*TreeNode
}

func newTreeNodeStack() *treeNodeStack {
	return &treeNodeStack{elements: []*TreeNode{}}
}

func (stack *treeNodeStack) push(element *TreeNode) {
	stack.elements = append(stack.elements, element)
}

func (stack *treeNodeStack) pop() *TreeNode {
	element := stack.get()
	if len(stack.elements) == 0 {
		return element
	}
	stack.elements = stack.elements[0 : len(stack.elements)-1]
	return element
}

func (stack *treeNodeStack) get() *TreeNode {
	length := len(stack.elements)
	if length == 0 {
		return nil
	}
	element := stack.elements[length-1]
	return element
}

func (stack *treeNodeStack) isEmpty() bool {
	return len(stack.elements) == 0
}
