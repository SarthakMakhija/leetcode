package add_2_numbers2

import (
	"reflect"
	"testing"
)

func TestAdd2Numbers1(t *testing.T) {
	l1 := &ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	newHead := addTwoNumbers(l1, l2)
	allValues := newHead.all()
	expected := []int{7, 8, 0, 7}

	if !reflect.DeepEqual(expected, allValues) {
		t.Fatalf("Expected %v, received %v", expected, allValues)
	}
}

func TestAdd2Numbers2(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	newHead := addTwoNumbers(l1, l2)
	allValues := newHead.all()
	expected := []int{8, 0, 7}

	if !reflect.DeepEqual(expected, allValues) {
		t.Fatalf("Expected %v received %v", expected, allValues)
	}
}

func TestAdd2Numbers3(t *testing.T) {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}

	newHead := addTwoNumbers(l1, l2)
	allValues := newHead.all()
	expected := []int{0}

	if !reflect.DeepEqual(expected, allValues) {
		t.Fatalf("Expected %v, received %v", expected, allValues)
	}
}

func TestAdd2Numbers4(t *testing.T) {
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{Val: 5}

	newHead := addTwoNumbers(l1, l2)
	allValues := newHead.all()
	expected := []int{1, 0}

	if !reflect.DeepEqual(expected, allValues) {
		t.Fatalf("Expected %v, received %v", expected, allValues)
	}
}
