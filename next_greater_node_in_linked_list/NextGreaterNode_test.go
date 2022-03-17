package next_greater_node_in_linked_list

import (
	"reflect"
	"testing"
)

func TestNextGreaterNodeValues1(t *testing.T) {
	head := &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 5}}}
	nextGreaterNodeValues := nextLargerNodes(head)
	expected := []int{5, 5, 0}

	if !reflect.DeepEqual(expected, nextGreaterNodeValues) {
		t.Fatalf("Expected %v, received %v", expected, nextGreaterNodeValues)
	}
}

func TestNextGreaterNodeValues2(t *testing.T) {
	head := &ListNode{Val: 2, Next: &ListNode{Val: 7, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}}
	nextGreaterNodeValues := nextLargerNodes(head)
	expected := []int{7, 0, 5, 5, 0}

	if !reflect.DeepEqual(expected, nextGreaterNodeValues) {
		t.Fatalf("Expected %v, received %v", expected, nextGreaterNodeValues)
	}
}
