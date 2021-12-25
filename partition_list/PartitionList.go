package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var nodes []*ListNode
	var partitionedListHead, partitionedListCurrent *ListNode = nil, nil

	listHead := head
	for listHead != nil {
		if listHead.Val >= x {
			nodes = append(nodes, listHead)
		} else {
			partitionedListHead, partitionedListCurrent = newNode(listHead, partitionedListHead, partitionedListCurrent)
		}
		listHead = listHead.Next
	}
	for _, node := range nodes {
		partitionedListHead, partitionedListCurrent = newNode(node, partitionedListHead, partitionedListCurrent)
	}

	return partitionedListHead
}

func newNode(node *ListNode, partitionedListHead *ListNode, partitionedListCurrent *ListNode) (*ListNode, *ListNode) {
	newNode := &ListNode{Val: node.Val}
	if partitionedListHead == nil {
		partitionedListHead, partitionedListCurrent = newNode, newNode
	} else {
		partitionedListCurrent.Next = newNode
		partitionedListCurrent = newNode
	}
	return partitionedListHead, partitionedListCurrent
}

func (listNode *ListNode) all() []int {
	var values []int
	head := listNode
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}
