package maximum_width_of_binary_tree

import (
	"testing"
)

func TestWidthOfBinaryTree1(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 9},
		},
	}
	width := widthOfBinaryTree(root)
	expected := 4

	if width != expected {
		t.Fatalf("Expected %v, received %v", expected, width)
	}
}

func TestWidthOfBinaryTree2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 9},
		},
	}
	width := widthOfBinaryTree(root)
	expected := 2

	if width != expected {
		t.Fatalf("Expected %v, received %v", expected, width)
	}
}

func TestWidthOfBinaryTree3(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}},
		},
		Right: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 7}},
		},
	}
	width := widthOfBinaryTree(root)
	expected := 7

	if width != expected {
		t.Fatalf("Expected %v, received %v", expected, width)
	}
}
