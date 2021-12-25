package next_right_pointer

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func (node *Node) isLeaf() bool {
	return node.Left == nil && node.Right == nil
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var connectInner func(node *Node)
	connectInner = func(node *Node) {
		/** It is a perfect binary tree, hence every node WILL have 2 children **/
		switch {
		case node == nil:
			return
		case node.isLeaf():
			return
		}
		node.Left.Next = node.Right
		if !node.isLeaf() && node.Right != nil && node.Next != nil {
			node.Right.Next = node.Next.Left
		}
		connectInner(node.Left)
		connectInner(node.Right)
	}
	connectInner(root)
	return root
}

func (node *Node) all() []int {
	values := []int{node.Val}
	head := node

	for head != nil {
		head = head.Left
		values = append(values, head.traverseUsingNext()...)
	}
	return values
}

func (node *Node) traverseUsingNext() []int {
	var values []int
	head := node

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}
