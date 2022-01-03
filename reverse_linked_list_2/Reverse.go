package reverse_linked_list_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var previous, rightmostNext *ListNode = nil, nil
	count, front := 1, head

	front, previous = move(front, previous, func() bool {
		count = count + 1
		return count <= left
	})
	newHead, backUp, rear := head, previous, front

	var reverseInner func(rear *ListNode, previous *ListNode, iterationCount int)
	reverseInner = func(rear *ListNode, previous *ListNode, iterationCount int) {
		if iterationCount >= right {
			rightmostNext = rear.Next
			rear.Next = previous
			if backUp != nil {
				backUp.Next = rear
			}
			if left == 1 {
				newHead = rear
			}
			return
		}
		if previous == nil {
			reverseInner(rear.Next, rear, iterationCount+1)
		} else {
			reverseInner(rear.Next, previous.Next, iterationCount+1)
		}
		rear.Next = previous
	}
	reverseInner(rear, previous, count-1)
	front.Next = rightmostNext

	return newHead
}

func move(ptr *ListNode, previous *ListNode, f func() bool) (*ListNode, *ListNode) {
	for f() {
		previous = ptr
		ptr = ptr.Next
	}
	return ptr, previous
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
