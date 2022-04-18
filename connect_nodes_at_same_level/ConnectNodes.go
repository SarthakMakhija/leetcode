package connect_nodes_at_same_level

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var connectInner func(node *Node, parent *Node)
	connectInner = func(node *Node, parent *Node) {
		if node == nil {
			return
		}
		//connect its children
		if node.Left != nil {
			node.Left.Next = node.Right
		}
		//attempt to connect the node
		if parent != nil {
			if node.Next == nil {
				if parent.Next != nil {
					nextOfParent := parent.Next
					for node.Next == nil && nextOfParent != nil {
						if nextOfParent.Left != nil {
							node.Next = nextOfParent.Left
						} else {
							node.Next = nextOfParent.Right
						}
						nextOfParent = nextOfParent.Next
					}
				}
			}
		}

		connectInner(node.Right, node)
		connectInner(node.Left, node)
	}
	connectInner(root, nil)
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
