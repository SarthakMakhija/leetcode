package delete_duplicates_from_sorted_List_2

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates1(t *testing.T) {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	head := deleteDuplicates(node)

	allValues := head.all()
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(expected, allValues) {
		t.Fatalf("Expected %v, received %v", expected, allValues)
	}

}

func TestRemoveDuplicates2(t *testing.T) {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}
	head := deleteDuplicates(node)

	allValues := head.all()
	expected := []int{1, 3}

	if !reflect.DeepEqual(expected, allValues) {
		t.Fatalf("Expected %v, received %v", expected, allValues)
	}
}

func TestRemoveDuplicates3(t *testing.T) {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}}
	head := deleteDuplicates(node)

	allValues := head.all()
	expected := []int{1}

	if !reflect.DeepEqual(expected, allValues) {
		t.Fatalf("Expected %v, received %v", expected, allValues)
	}
}

func TestRemoveDuplicates4(t *testing.T) {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}
	head := deleteDuplicates(node)

	allValues := head.all()
	expected := []int{2}

	if !reflect.DeepEqual(expected, allValues) {
		t.Fatalf("Expected %v, received %v", expected, allValues)
	}
}

func TestRemoveDuplicates5(t *testing.T) {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}
	head := deleteDuplicates(node)

	allValues := head.all()
	var expected []int

	if !reflect.DeepEqual(expected, allValues) {
		t.Fatalf("Expected %v, received %v", expected, allValues)
	}

}
