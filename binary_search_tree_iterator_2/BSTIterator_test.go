package binary_search_tree_iterator_2

import (
	"testing"
)

func TestBSTIterator1(t *testing.T) {
	tree := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}

	expectedValues, expectedBools := []int{3, 7, 9, 15, 20}, []bool{true, true, true, true, true}
	iterator := Constructor(tree)

	for count := 1; count <= 5; count++ {
		if hasNext := iterator.HasNext(); hasNext != expectedBools[count-1] {
			t.Fatalf("Expected %v, received %v", expectedBools[count-1], hasNext)
		}
		if value := iterator.Next(); value != expectedValues[count-1] {
			t.Fatalf("Expected %v, received %v", expectedValues[count-1], value)
		}
	}
	if hasNext := iterator.HasNext(); hasNext != false {
		t.Fatalf("Expected %v, received %v", false, hasNext)
	}
}

func TestBSTIterator2(t *testing.T) {
	tree := &TreeNode{Val: 7,
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}

	iterator := Constructor(tree)
	expectedValues, expectedBools := []int{7, 9, 15, 20}, []bool{true, true, true, true}

	for count := 1; count <= 4; count++ {
		if hasNext := iterator.HasNext(); hasNext != expectedBools[count-1] {
			t.Fatalf("Expected %v, received %v", expectedBools[count-1], hasNext)
		}
		if value := iterator.Next(); value != expectedValues[count-1] {
			t.Fatalf("Expected %v, received %v", expectedValues[count-1], value)
		}
	}
	if hasNext := iterator.HasNext(); hasNext != false {
		t.Fatalf("Expected %v, received %v", false, hasNext)
	}
}

func TestBSTIterator3(t *testing.T) {
	tree := &TreeNode{Val: 7,
		Left:  &TreeNode{Val: 3, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}}

	iterator := Constructor(tree)
	expectedValues, expectedBools := []int{3, 5, 7, 9, 15, 20}, []bool{true, true, true, true, true, true}

	for count := 1; count <= 6; count++ {
		if hasNext := iterator.HasNext(); hasNext != expectedBools[count-1] {
			t.Fatalf("Expected %v, received %v", expectedBools[count-1], hasNext)
		}
		if value := iterator.Next(); value != expectedValues[count-1] {
			t.Fatalf("Expected %v, received %v", expectedValues[count-1], value)
		}
	}
	if hasNext := iterator.HasNext(); hasNext != false {
		t.Fatalf("Expected %v, received %v", false, hasNext)
	}
}
