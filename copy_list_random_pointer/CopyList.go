package copy_list_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newNodeAddressByOldNodeAddress := make(map[*Node]*Node)

	newHead := &Node{Val: head.Val}
	newCurrent := newHead
	newNodeAddressByOldNodeAddress[head] = newHead

	for head != nil {
		if head.Next != nil {
			node, ok := newNodeAddressByOldNodeAddress[(head.Next)]
			if ok {
				newCurrent.Next = node
			} else {
				newCurrent.Next = &Node{Val: head.Next.Val}
				newNodeAddressByOldNodeAddress[(head.Next)] = newCurrent.Next
			}
		}
		if head.Random != nil {
			node, ok := newNodeAddressByOldNodeAddress[(head.Random)]
			if ok {
				newCurrent.Random = node
			} else {
				newCurrent.Random = &Node{Val: head.Random.Val}
				newNodeAddressByOldNodeAddress[(head.Random)] = newCurrent.Random
			}
		}
		newCurrent = newCurrent.Next
		head = head.Next
	}
	return newHead
}

func (node *Node) all() []ListWithRandomPointerValue {
	head := node
	var values []ListWithRandomPointerValue

	for head != nil {
		randomValue := -1
		if head.Random != nil {
			randomValue = head.Random.Val
		}
		values = append(values, ListWithRandomPointerValue{Val: head.Val, RandomVal: randomValue})
		head = head.Next
	}
	return values
}

type ListWithRandomPointerValue struct {
	Val, RandomVal int
}
