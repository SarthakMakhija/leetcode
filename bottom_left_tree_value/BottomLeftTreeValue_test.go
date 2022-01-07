package bottom_left_tree_value

import (
	"testing"
)

func TestBottomLeftValue1(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 6}},
	}

	bottomLeft := findBottomLeftValue(root)
	if bottomLeft != 7 {
		t.Fatalf("Expected %v, received %v", 7, bottomLeft)
	}
}

func TestBottomLeftValue2(t *testing.T) {
	root := &TreeNode{
		Val:   0,
		Left:  &TreeNode{Val: -1},
		Right: nil,
	}

	bottomLeft := findBottomLeftValue(root)
	if bottomLeft != -1 {
		t.Fatalf("Expected %v, received %v", -1, bottomLeft)
	}
}
