package delete_duplicates_from_sorted_List_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var previous *ListNode = nil
	var newHead *ListNode = nil

	front, rear := head, head
	for front != nil && rear != nil {
		rear = move(rear, func(ptr *ListNode) bool {
			return ptr != nil && front.Val == ptr.Val
		})
		if front.Next == rear {
			previous = front
			if newHead == nil {
				newHead = front
			}
		} else {
			if previous != nil {
				previous.Next = rear
			}
		}
		if rear != nil {
			front = rear
		}
	}
	return newHead
}

func move(ptr *ListNode, f func(ptr *ListNode) bool) *ListNode {
	for f(ptr) {
		ptr = ptr.Next
	}
	return ptr
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
