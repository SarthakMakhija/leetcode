package add_2_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(first *ListNode, second *ListNode) *ListNode {
	var resultListHead, resultListCurrent *ListNode = nil, nil
	firstHead, secondHead, carry := first, second, 0

	for firstHead != nil && secondHead != nil {
		sum := firstHead.Val + secondHead.Val + carry
		carry = sum / 10
		digit := sum % 10

		node := &ListNode{Val: digit, Next: nil}
		if resultListHead == nil || resultListCurrent == nil {
			resultListHead, resultListCurrent = node, node
		} else {
			resultListCurrent.Next, resultListCurrent = node, node
		}
		firstHead, secondHead = firstHead.Next, secondHead.Next
	}

	if firstHead != nil {
		addDigits(firstHead, resultListCurrent, carry)
	} else if secondHead != nil {
		addDigits(secondHead, resultListCurrent, carry)
	} else {
		addCarryNode(carry, resultListCurrent)
	}
	return resultListHead
}

func addDigits(head *ListNode, resultListCurrent *ListNode, carry int) {
	for head != nil {
		sum := head.Val + carry
		carry = sum / 10
		digit := sum % 10
		node := &ListNode{Val: digit, Next: nil}
		resultListCurrent.Next, resultListCurrent = node, node
		head = head.Next
	}
	addCarryNode(carry, resultListCurrent)
}

func addCarryNode(carry int, resultListCurrent *ListNode) {
	if carry != 0 {
		node := &ListNode{Val: carry, Next: nil}
		resultListCurrent.Next, resultListCurrent = node, node
	}
}

func (list *ListNode) allValues() []int {
	var values []int
	head := list

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}
