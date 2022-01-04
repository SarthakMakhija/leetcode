package delete_node_in_BST

import (
	"testing"
)

func TestDeleteNode1(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}

	newRoot := deleteNode(root, 4)
	values := newRoot.traverse()
	expected := "23567"

	if expected != values {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestDeleteNode2(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}

	newRoot := deleteNode(root, 7)
	values := newRoot.traverse()
	expected := "23456"

	if expected != values {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestDeleteNode3(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}

	newRoot := deleteNode(root, 3)
	values := newRoot.traverse()
	expected := "24567"

	if expected != values {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestDeleteNode4(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}

	newRoot := deleteNode(root, 5)
	values := newRoot.traverse()
	expected := "23467"

	if expected != values {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestDeleteNode5(t *testing.T) {
	root := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}

	newRoot := deleteNode(root, 0)
	values := newRoot.traverse()
	expected := "234567"

	if expected != values {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestDeleteNode6(t *testing.T) {
	root := &TreeNode{Val: 0}

	newRoot := deleteNode(root, 0)
	values := newRoot.traverse()
	expected := ""

	if expected != values {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestDeleteNode7(t *testing.T) {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	newRoot := deleteNode(root, 1)
	values := newRoot.traverse()
	expected := "2"

	if expected != values {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}
