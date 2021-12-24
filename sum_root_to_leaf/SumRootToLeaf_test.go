package sum_root_to_leaf

import (
	"testing"
)

func TestSumRootToLeaf1(t *testing.T) {
	tree := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}}},
		Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}}

	sum := sumNumbers(tree)
	expected := 10382

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestSumRootToLeaf2(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3}}

	sum := sumNumbers(tree)
	expected := 25

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestSumRootToLeaf3(t *testing.T) {
	tree := &TreeNode{Val: 1}

	sum := sumNumbers(tree)
	expected := 1

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}

func TestSumRootToLeaf4(t *testing.T) {
	tree := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 0}}

	sum := sumNumbers(tree)
	expected := 1026

	if sum != expected {
		t.Fatalf("Expected %v, received %v", expected, sum)
	}
}
