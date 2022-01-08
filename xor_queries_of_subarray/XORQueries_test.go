package xor_queries_of_subarray

import (
	"reflect"
	"testing"
)

func TestXorQueries1(t *testing.T) {
	replies := xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}})
	expected := []int{2}

	if !reflect.DeepEqual(expected, replies) {
		t.Fatalf("Expected %v, received %v", expected, replies)
	}
}

func TestXorQueries2(t *testing.T) {
	replies := xorQueries([]int{1, 3, 4, 8}, [][]int{{1, 2}})
	expected := []int{7}

	if !reflect.DeepEqual(expected, replies) {
		t.Fatalf("Expected %v, received %v", expected, replies)
	}
}

func TestXorQueries3(t *testing.T) {
	replies := xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 3}})
	expected := []int{14}

	if !reflect.DeepEqual(expected, replies) {
		t.Fatalf("Expected %v, received %v", expected, replies)
	}
}

func TestXorQueries4(t *testing.T) {
	replies := xorQueries([]int{1, 3, 4, 8}, [][]int{{3, 3}})
	expected := []int{8}

	if !reflect.DeepEqual(expected, replies) {
		t.Fatalf("Expected %v, received %v", expected, replies)
	}
}
