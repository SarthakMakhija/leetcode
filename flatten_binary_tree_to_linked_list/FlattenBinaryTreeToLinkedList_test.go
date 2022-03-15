package flatten_binary_tree_to_linked_list

import (
	"reflect"
	"testing"
)

func TestFlatten1(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
	}
	flatten(root)
	values := travelRight(root)
	expected := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestFlatten2(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}},
	}
	flatten(root)
	values := travelRight(root)
	expected := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestFlatten3(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 40}}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}},
	}
	flatten(root)
	values := travelRight(root)
	expected := []int{1, 2, 3, 4, 40, 5, 6, 7}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestFlatten4(t *testing.T) {
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 40}}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}},
	}
	flatten(root)
	values := travelRight(root)
	expected := []int{1, 2, 3, 4, 40, 5, 6, 7}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}
