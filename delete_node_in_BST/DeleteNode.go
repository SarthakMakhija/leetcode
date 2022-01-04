package delete_node_in_BST

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	node, parent := search(root, key)
	if node == nil {
		return root
	}
	if isLeaf(node) {
		if parent == nil {
			return nil
		}
		deleteLeaf(node, parent)
		return root
	}
	deleteInternal(node)
	return root
}

func search(node *TreeNode, key int) (*TreeNode, *TreeNode) {
	var searchInner func(node, previous *TreeNode) (*TreeNode, *TreeNode)
	searchInner = func(node, previous *TreeNode) (*TreeNode, *TreeNode) {
		if node == nil {
			return nil, nil
		} else if key > node.Val {
			return searchInner(node.Right, node)
		} else if key < node.Val {
			return searchInner(node.Left, node)
		} else {
			return node, previous
		}
	}
	return searchInner(node, nil)
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}

func deleteLeaf(node *TreeNode, parent *TreeNode) {
	if parent == nil {
		node = nil
		return
	}
	if parent.Left == node {
		parent.Left = nil
	} else {
		parent.Right = nil
	}
}

func deleteInternal(node *TreeNode) {
	predecessor, parent, isLeaf := inOrderPredecessor(node)
	if predecessor != nil {
		node.Val = predecessor.Val
		if isLeaf {
			deleteLeaf(predecessor, parent)
		} else {
			deleteInternal(predecessor)
		}
	} else {
		successor, parent, isLeaf := inOrderSuccessor(node)
		if successor != nil {
			node.Val = successor.Val
			if isLeaf {
				deleteLeaf(successor, parent)
			} else {
				deleteInternal(successor)
			}
		}
	}
}

func inOrderPredecessor(node *TreeNode) (*TreeNode, *TreeNode, bool) {
	parent, n := node, node.Left
	if n != nil {
		for n.Right != nil {
			parent = n
			n = n.Right
		}
	}
	return n, parent, isLeaf(n)
}

func inOrderSuccessor(node *TreeNode) (*TreeNode, *TreeNode, bool) {
	parent, n := node, node.Right
	if n != nil {
		for n.Left != nil {
			parent = n
			n = n.Left
		}
	}
	return n, parent, isLeaf(n)
}

func (tree *TreeNode) traverse() string {
	var traverseInner func(node *TreeNode) string
	traverseInner = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		return traverseInner(node.Left) + strconv.Itoa(node.Val) + traverseInner(node.Right)
	}
	return traverseInner(tree)
}
