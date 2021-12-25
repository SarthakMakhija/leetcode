package partition_list

import (
	"reflect"
	"testing"
)

func TestPartition1(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}

	partitioned := partition(list, 3)
	received := partitioned.all()
	expected := []int{1, 2, 2, 4, 3, 5}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected %v , received %v", expected, received)
	}
}

func TestPartition2(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}

	partitioned := partition(list, 1)
	received := partitioned.all()
	expected := []int{1, 4, 3, 2, 5, 2}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected %v , received %v", expected, received)
	}
}

func TestPartition3(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}

	partitioned := partition(list, 6)
	received := partitioned.all()
	expected := []int{1, 4, 3, 2, 5, 2}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected %v , received %v", expected, received)
	}
}

func TestPartition4(t *testing.T) {
	list := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}

	partitioned := partition(list, 5)
	received := partitioned.all()
	expected := []int{1, 4, 3, 2, 2, 5}

	if !reflect.DeepEqual(expected, received) {
		t.Fatalf("Expected %v , received %v", expected, received)
	}
}
