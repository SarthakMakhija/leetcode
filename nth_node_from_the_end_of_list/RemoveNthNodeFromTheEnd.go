package nth_node_from_the_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var tail *ListNode = nil
	first := head

	firstJumpCount := 0
	for first.Next != nil {
		first = first.Next
		firstJumpCount = firstJumpCount + 1
	}
	tailJumpCount := firstJumpCount - n + 1
	for jump := 1; jump <= tailJumpCount; jump++ {
		if tail == nil {
			tail = head
		} else {
			tail = tail.Next
		}
	}

	if tail == nil {
		tail, head = head, head.Next
		tail.Next = nil
		return head
	}
	originalNext := tail.Next
	tail.Next = tail.Next.Next
	if originalNext != nil {
		originalNext.Next = nil
	}
	return head
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
