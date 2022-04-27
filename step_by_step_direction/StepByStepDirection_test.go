package step_by_step_direction

import (
	"testing"
)

func TestStepByStepDirection1(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 4}},
	}
	directions := getDirections(root, 3, 6)
	expected := "UURL"

	if directions != expected {
		t.Fatalf("Expected %v, received %v", expected, directions)
	}
}

func TestStepByStepDirection2(t *testing.T) {
	root := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}},
	}
	directions := getDirections(root, 3, 6)
	expected := "UR"

	if directions != expected {
		t.Fatalf("Expected %v, received %v", expected, directions)
	}
}

func TestStepByStepDirection3(t *testing.T) {
	root := &TreeNode{
		Val:  2,
		Left: &TreeNode{Val: 1},
	}
	directions := getDirections(root, 2, 1)
	expected := "L"

	if directions != expected {
		t.Fatalf("Expected %v, received %v", expected, directions)
	}
}

func TestStepByStepDirection4(t *testing.T) {
	root := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}}}},
	}
	directions := getDirections(root, 3, 6)
	expected := "RL"

	if directions != expected {
		t.Fatalf("Expected %v, received %v", expected, directions)
	}
}

func TestStepByStepDirection5(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	directions := getDirections(root, 2, 1)
	expected := "U"

	if directions != expected {
		t.Fatalf("Expected %v, received %v", expected, directions)
	}
}

func TestStepByStepDirection6(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 4}},
	}
	directions := getDirections(root, 1, 6)
	expected := "URL"

	if directions != expected {
		t.Fatalf("Expected %v, received %v", expected, directions)
	}
}
