package add_row_to_tree

import (
	"testing"
)

func TestAddOneRow1(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}}
	addOneRow(root, 1, 3)

	expected := "312114"
	values := root.traverse()

	if expected != values {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestAddOneRow2(t *testing.T) {
	root := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}},
	}

	addOneRow(root, 1, 2)

	expected := "32114156"
	values := root.traverse()

	if expected != values {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}
