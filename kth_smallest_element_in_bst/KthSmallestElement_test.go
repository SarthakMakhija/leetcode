package kth_smallest_element_in_bst

import (
	"testing"
)

func TestKthSmallestElement1(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	smallest := kthSmallest(root, 1)

	if smallest != 1 {
		t.Fatalf("Expected %vst smallest to be %v, received %v", 1, 1, smallest)
	}
}

func TestKthSmallestElement2(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	smallest := kthSmallest(root, 2)

	if smallest != 2 {
		t.Fatalf("Expected %vnd smallest to be %v, received %v", 2, 2, smallest)
	}
}

func TestKthSmallestElement3(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	smallest := kthSmallest(root, 3)

	if smallest != 3 {
		t.Fatalf("Expected %vrd smallest to be %v, received %v", 3, 3, smallest)
	}
}

func TestKthSmallestElement4(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	smallest := kthSmallest(root, 4)

	if smallest != 4 {
		t.Fatalf("Expected %vrd smallest to be %v, received %v", 4, 4, smallest)
	}
}
