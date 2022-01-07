package largest_value_in_each_tree_row

import (
	"reflect"
	"testing"
)

func TestLargestValueInEachTreeRow1(t *testing.T) {
	tree := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 9}}}

	largestValues := largestValues(tree)
	expected := []int{1, 3, 9}

	if !reflect.DeepEqual(expected, largestValues) {
		t.Fatalf("Expected %v, received %v", expected, largestValues)
	}
}

func TestLargestValueInEachTreeRow2(t *testing.T) {
	tree := &TreeNode{Val: 0}

	largestValues := largestValues(tree)
	expected := []int{0}

	if !reflect.DeepEqual(expected, largestValues) {
		t.Fatalf("Expected %v, received %v", expected, largestValues)
	}
}

func TestLargestValueInEachTreeRow3(t *testing.T) {
	tree := &TreeNode{Val: 0, Left: &TreeNode{Val: -1}}

	largestValues := largestValues(tree)
	expected := []int{0, -1}

	if !reflect.DeepEqual(expected, largestValues) {
		t.Fatalf("Expected %v, received %v", expected, largestValues)
	}
}
