package top_k_frequent_elements

import (
	"reflect"
	"testing"
)

func TestTopKFrequentElements1(t *testing.T) {
	topK := topKFrequent([]int{3, 0, 1, 0}, 1)
	expected := []int{0}

	if !reflect.DeepEqual(expected, topK) {
		t.Fatalf("Expected %v, received %v", expected, topK)
	}
}

func TestTopKFrequentElements2(t *testing.T) {
	topK := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	expected := []int{1, 2}

	if !reflect.DeepEqual(expected, topK) {
		t.Fatalf("Expected %v, received %v", expected, topK)
	}
}
