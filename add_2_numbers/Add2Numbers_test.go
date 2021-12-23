package add_2_numbers

import (
	"reflect"
	"testing"
)

func TestAdd2NumbersWithSameNumberOfDigits1(t *testing.T) {
	result := addTwoNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}})

	sum := result.allValues()
	expected := []int{7, 0, 8}

	if !reflect.DeepEqual(expected, sum) {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestAdd2NumbersWithSameNumberOfDigits2(t *testing.T) {
	result := addTwoNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}})

	sum := result.allValues()
	expected := []int{4, 8, 6}

	if !reflect.DeepEqual(expected, sum) {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestAdd2NumbersWithSameNumberOfDigits3(t *testing.T) {
	result := addTwoNumbers(
		&ListNode{Val: 5},
		&ListNode{Val: 5})

	sum := result.allValues()
	expected := []int{0, 1}

	if !reflect.DeepEqual(expected, sum) {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestAdd2NumbersWithDifferentNumberOfDigits1(t *testing.T) {
	result := addTwoNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4}},
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}})

	sum := result.allValues()
	expected := []int{4, 8, 3}

	if !reflect.DeepEqual(expected, sum) {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestAdd2NumbersWithDifferentNumberOfDigits2(t *testing.T) {
	result := addTwoNumbers(
		&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}},
		&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}})

	sum := result.allValues()
	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

	if !reflect.DeepEqual(expected, sum) {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}
