package right_view_of_binary_tree

import (
	"reflect"
	"testing"
)

func TestRightView1(t *testing.T) {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
	}
	rightView := rightSideView(root)
	expected := []int{1, 3, 4}

	if !reflect.DeepEqual(expected, rightView) {
		t.Fatalf("Expected %v, received %v", expected, rightView)
	}
}

func TestRightView2(t *testing.T) {
	root := &TreeNode{
		Val:  10,
		Left: &TreeNode{Val: 20, Right: &TreeNode{Val: 5}},
	}
	rightView := rightSideView(root)
	expected := []int{10, 20, 5}

	if !reflect.DeepEqual(expected, rightView) {
		t.Fatalf("Expected %v, received %v", expected, rightView)
	}
}
