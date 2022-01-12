package positive_negative_pair

import (
	"reflect"
	"testing"
)

func TestFindPairs1(t *testing.T) {
	arr := []int{1, 3, 6, -2, -1, -3, 2, 7}
	pairs := findPairs(arr)
	expected := []int{-1, 1, -3, 3, -2, 2}

	if !reflect.DeepEqual(expected, pairs) {
		t.Fatalf("Expectede %v, received %v", expected, pairs)
	}
}

func TestFindPairs2(t *testing.T) {
	arr := []int{3, 2, 1}
	pairs := findPairs(arr)
	var expected []int

	if !reflect.DeepEqual(expected, pairs) {
		t.Fatalf("Expectede %v, received %v", expected, pairs)
	}
}

func TestFindPairs3(t *testing.T) {
	arr := []int{-1, 2, -1, -2}
	pairs := findPairs(arr)
	expected := []int{-2, 2}

	if !reflect.DeepEqual(expected, pairs) {
		t.Fatalf("Expectede %v, received %v", expected, pairs)
	}
}

func TestFindPairs4(t *testing.T) {
	arr := []int{-1, 2, 1, -1, -2}
	pairs := findPairs(arr)
	expected := []int{-1, 1, -2, 2}

	if !reflect.DeepEqual(expected, pairs) {
		t.Fatalf("Expectede %v, received %v", expected, pairs)
	}
}

func TestFindPairs5(t *testing.T) {
	arr := []int{-1, 2, 1, -1, 1, -2}
	pairs := findPairs(arr)
	expected := []int{-1, 1, -1, 1, -2, 2}

	if !reflect.DeepEqual(expected, pairs) {
		t.Fatalf("Expectede %v, received %v", expected, pairs)
	}
}
