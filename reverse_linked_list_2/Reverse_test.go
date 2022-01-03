package reverse_linked_list_2

import (
	"reflect"
	"testing"
)

func TestReverseList1(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head := reverseBetween(list, 2, 4)

	values := head.all()
	expected := []int{1, 4, 3, 2, 5}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestReverseList2(t *testing.T) {
	list := &ListNode{Val: 1}
	head := reverseBetween(list, 1, 1)

	values := head.all()
	expected := []int{1}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestReverseList3(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head := reverseBetween(list, 1, 3)

	values := head.all()
	expected := []int{3, 2, 1, 4, 5}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}
