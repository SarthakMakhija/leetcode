package add_row_to_tree

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		newRoot := &TreeNode{Val: val}
		newRoot.Left = root
		return newRoot
	}
	var addOneInner func(node *TreeNode, currentDepth int)
	addOneInner = func(node *TreeNode, currentDepth int) {
		if node == nil {
			return
		}
		if currentDepth == depth-1 {
			newLeft, newRight := &TreeNode{Val: val}, &TreeNode{Val: val}
			newLeft.Left, newRight.Right = node.Left, node.Right
			node.Left, node.Right = newLeft, newRight
			return
		}
		addOneInner(node.Left, currentDepth+1)
		addOneInner(node.Right, currentDepth+1)
	}
	addOneInner(root, 1)
	return root
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
