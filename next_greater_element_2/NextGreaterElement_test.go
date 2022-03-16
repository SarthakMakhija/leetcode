package next_greater_element_2

import (
	"reflect"
	"testing"
)

func TestNextGreater1(t *testing.T) {
	nextGreaterElements := nextGreaterElements([]int{1, 2, 3, 4})
	expected := []int{2, 3, 4, -1}

	if !reflect.DeepEqual(expected, nextGreaterElements) {
		t.Fatalf("Expected %v, received %v", expected, nextGreaterElements)
	}
}

func TestNextGreater2(t *testing.T) {
	nextGreaterElements := nextGreaterElements([]int{1, 2, 3, 0})
	expected := []int{2, 3, -1, 1}

	if !reflect.DeepEqual(expected, nextGreaterElements) {
		t.Fatalf("Expected %v, received %v", expected, nextGreaterElements)
	}
}

func TestNextGreater3(t *testing.T) {
	nextGreaterElements := nextGreaterElements([]int{10, 2, 1, 6, 9})
	expected := []int{-1, 6, 6, 9, 10}

	if !reflect.DeepEqual(expected, nextGreaterElements) {
		t.Fatalf("Expected %v, received %v", expected, nextGreaterElements)
	}
}

func TestNextGreater4(t *testing.T) {
	nextGreaterElements := nextGreaterElements([]int{10, 2, 1, 6, 6})
	expected := []int{-1, 6, 6, 10, 10}

	if !reflect.DeepEqual(expected, nextGreaterElements) {
		t.Fatalf("Expected %v, received %v", expected, nextGreaterElements)
	}
}
