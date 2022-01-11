package count_distinct_elements_in_each_window

import (
	"reflect"
	"testing"
)

func TestCountDistinct1(t *testing.T) {
	nums := []int{10, 20, 20, 10, 30, 40, 10}
	distinct := countDistinct(nums, 4)
	expected := []int{2, 3, 4, 3}

	if !reflect.DeepEqual(expected, distinct) {
		t.Fatalf("Expected %v, received %v", expected, distinct)
	}
}

func TestCountDistinct2(t *testing.T) {
	nums := []int{10, 10, 10, 10}
	distinct := countDistinct(nums, 3)
	expected := []int{1, 1}

	if !reflect.DeepEqual(expected, distinct) {
		t.Fatalf("Expected %v, received %v", expected, distinct)
	}
}
