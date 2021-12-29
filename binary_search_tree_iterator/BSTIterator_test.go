package binary_search_tree_iterator

import (
	"reflect"
	"testing"
)

func TestBSTIterator1(t *testing.T) {
	tree := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}

	iterator := Constructor(tree)
	next := iterator.Next()
	expected := 3

	if next != expected {
		t.Fatalf("Expected %v, received %v", expected, next)
	}
}

func TestBSTIterator2(t *testing.T) {
	tree := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}

	iterator := Constructor(tree)
	var values []int
	values = append(values, iterator.Next(), iterator.Next(), iterator.Next(), iterator.Next(), iterator.Next())
	expected := []int{3, 7, 9, 15, 20}

	if !reflect.DeepEqual(expected, values) {
		t.Fatalf("Expected %v, received %v", expected, values)
	}
}

func TestBSTIterator3(t *testing.T) {
	tree := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}

	iterator := Constructor(tree)
	iterator.Next()
	iterator.Next()
	iterator.Next()
	iterator.Next()
	iterator.Next()

	hasNext := iterator.HasNext()

	if hasNext != false {
		t.Fatalf("Expected %v, received %v", false, hasNext)
	}
}
