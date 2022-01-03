package majority_element

import (
	"reflect"
	"testing"
)

func TestMajority1(t *testing.T) {
	majority := majorityElement([]int{3, 2, 3})
	expected := []int{3}

	if !reflect.DeepEqual(expected, majority) {
		t.Fatalf("Expected %v, received %v", expected, majority)
	}
}

func TestMajority2(t *testing.T) {
	majority := majorityElement([]int{1})
	expected := []int{1}

	if !reflect.DeepEqual(expected, majority) {
		t.Fatalf("Expected %v, received %v", expected, majority)
	}
}

func TestMajority3(t *testing.T) {
	majority := majorityElement([]int{1, 2})
	expected := []int{1, 2}

	if !reflect.DeepEqual(expected, majority) {
		t.Fatalf("Expected %v, received %v", expected, majority)
	}
}

func TestMajority4(t *testing.T) {
	majority := majorityElement([]int{2, 2})
	expected := []int{2}

	if !reflect.DeepEqual(expected, majority) {
		t.Fatalf("Expected %v, received %v", expected, majority)
	}
}

func TestMajority5(t *testing.T) {
	majority := majorityElement([]int{1, 2, 3, 5, 6, 7})
	var expected []int

	if !reflect.DeepEqual(expected, majority) {
		t.Fatalf("Expected %v, received %v", expected, majority)
	}
}

func TestMajority6(t *testing.T) {
	majority := majorityElement([]int{1, 2, 3, 5, 6, 7, 7, 7})
	expected := []int{7}

	if !reflect.DeepEqual(expected, majority) {
		t.Fatalf("Expected %v, received %v", expected, majority)
	}
}

func TestMajority7(t *testing.T) {
	majority := majorityElement([]int{1, 2, 3, 3, 3, 5, 3, 3, 3, 3, 3})
	expected := []int{3}

	if !reflect.DeepEqual(expected, majority) {
		t.Fatalf("Expected %v, received %v", expected, majority)
	}
}
