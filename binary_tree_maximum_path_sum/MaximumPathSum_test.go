package binary_tree_maximum_path_sum

import (
	"testing"
)

func TestMaxPathSum1(t *testing.T) {
	root := &TreeNode{
		Val:   -10,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	maxSum := maxPathSum(root)
	expected := 42

	if maxSum != expected {
		t.Fatalf("Expected %v, received %v", expected, maxSum)
	}
}

func TestMaxPathSum2(t *testing.T) {
	root := &TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}
	maxSum := maxPathSum(root)
	expected := 54

	if maxSum != expected {
		t.Fatalf("Expected %v, received %v", expected, maxSum)
	}
}

func TestMaxPathSum3(t *testing.T) {
	root := &TreeNode{Val: -3}

	maxSum := maxPathSum(root)
	expected := -3

	if maxSum != expected {
		t.Fatalf("Expected %v, received %v", expected, maxSum)
	}
}

func TestMaxPathSum4(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: -2}, Right: &TreeNode{Val: 3}}

	maxSum := maxPathSum(root)
	expected := 4

	if maxSum != expected {
		t.Fatalf("Expected %v, received %v", expected, maxSum)
	}
}
