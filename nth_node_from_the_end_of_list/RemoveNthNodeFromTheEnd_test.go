package nth_node_from_the_end_of_list

import (
	"reflect"
	"testing"
)

func TestRemoveNthNodeFromList1(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4,
		Next: &ListNode{Val: 5}}}}}

	head := removeNthFromEnd(list, 2)
	values := head.all()
	expected := []int{1, 2, 3, 5}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestRemoveNthNodeFromList2(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4,
		Next: &ListNode{Val: 5}}}}}

	head := removeNthFromEnd(list, 3)
	values := head.all()
	expected := []int{1, 2, 4, 5}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestRemoveNthNodeFromList3(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4,
		Next: &ListNode{Val: 5}}}}}

	head := removeNthFromEnd(list, 1)
	values := head.all()
	expected := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestRemoveNthNodeFromList4(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2}}

	head := removeNthFromEnd(list, 1)
	values := head.all()
	expected := []int{1}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestRemoveNthNodeFromList5(t *testing.T) {
	list := &ListNode{Val: 1}

	head := removeNthFromEnd(list, 1)
	values := head.all()
	var expected []int

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestRemoveNthNodeFromList6(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4,
		Next: &ListNode{Val: 5}}}}}

	head := removeNthFromEnd(list, 5)
	values := head.all()
	expected := []int{2, 3, 4, 5}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}
