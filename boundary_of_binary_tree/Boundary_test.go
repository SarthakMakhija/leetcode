package boundary_of_binary_tree

import (
	"reflect"
	"testing"
)

func TestBoundaryOfABinaryTree1(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 8}},
		},
		Right: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 6, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 10}},
		},
	}
	boundary := boundaryOfBinaryTree(root)
	expected := []int{1, 2, 4, 7, 8, 9, 10, 6, 3}

	if !reflect.DeepEqual(expected, boundary) {
		t.Fatalf("Expected %v, received %v", expected, boundary)
	}
}

func TestBoundaryOfABinaryTree2(t *testing.T) {
	root := &TreeNode{
		Val: 1,

		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
	}
	boundary := boundaryOfBinaryTree(root)
	expected := []int{1, 3, 4, 2}

	if !reflect.DeepEqual(expected, boundary) {
		t.Fatalf("Expected %v, received %v", expected, boundary)
	}
}

func TestBoundaryOfABinaryTree3(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	boundary := boundaryOfBinaryTree(root)
	expected := []int{3, 5, 6, 7, 4, 0, 8, 1}

	if !reflect.DeepEqual(expected, boundary) {
		t.Fatalf("Expected %v, received %v", expected, boundary)
	}
}

func TestBoundaryOfABinaryTree4(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 4},
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 8},
		},
	}
	boundary := boundaryOfBinaryTree(root)
	expected := []int{3, 5, 2, 7, 4, 0, 8, 1}

	if !reflect.DeepEqual(expected, boundary) {
		t.Fatalf("Expected %v, received %v", expected, boundary)
	}
}
