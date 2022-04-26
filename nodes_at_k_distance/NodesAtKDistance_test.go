package nodes_at_k_distance

import (
	"reflect"
	"testing"
)

func TestNodesAtDistanceK1(t *testing.T) {
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
	nodes := distanceK(root, &TreeNode{Val: 5}, 2)
	expected := []int{7, 4, 1}

	if !reflect.DeepEqual(expected, nodes) {
		t.Fatalf("Expected %v, received %v", expected, nodes)
	}
}

func TestNodesAtDistanceK2(t *testing.T) {
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
	nodes := distanceK(root, &TreeNode{Val: 5}, 3)
	expected := []int{0, 8}

	if !reflect.DeepEqual(expected, nodes) {
		t.Fatalf("Expected %v, received %v", expected, nodes)
	}
}

func TestNodesAtDistanceK3(t *testing.T) {
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
	nodes := distanceK(root, &TreeNode{Val: 5}, 1)
	expected := []int{6, 2, 3}

	if !reflect.DeepEqual(expected, nodes) {
		t.Fatalf("Expected %v, received %v", expected, nodes)
	}
}

func TestNodesAtDistanceK4(t *testing.T) {
	root := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3},
		},
	}

	nodes := distanceK(root, &TreeNode{Val: 3}, 3)
	expected := []int{2}

	if !reflect.DeepEqual(expected, nodes) {
		t.Fatalf("Expected %v, received %v", expected, nodes)
	}
}
