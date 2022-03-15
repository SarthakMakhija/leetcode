package flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLeaf(treeNode *TreeNode) bool {
	return treeNode.Left == nil && treeNode.Right == nil
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var flattenInner func(node *TreeNode)
	flattenInner = func(node *TreeNode) {
		if node == nil {
			return
		}
		flattenInner(node.Left)
		flattenInner(node.Right)
		if !isLeaf(node) && node.Left != nil {
			originalRight := node.Right
			node.Right = node.Left
			node.Left = nil
			move := node
			for move.Right != nil {
				move = move.Right
			}
			move.Right = originalRight
		}
	}
	flattenInner(root)
}

func travelRight(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var travelRightInner func(node *TreeNode)
	var values []int

	travelRightInner = func(node *TreeNode) {
		if node == nil {
			return
		}
		values = append(values, node.Val)
		travelRightInner(node.Right)
	}
	travelRightInner(root)
	return values
}
