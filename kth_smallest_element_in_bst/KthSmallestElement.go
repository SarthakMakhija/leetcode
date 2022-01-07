package kth_smallest_element_in_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var kthSmallestInner func(node *TreeNode)
	index, value := 0, -1
	kthSmallestInner = func(node *TreeNode) {
		switch {
		case node == nil:
			return
		}
		kthSmallestInner(node.Left)
		index = index + 1

		if index == k {
			value = node.Val
			return
		}
		kthSmallestInner(node.Right)
	}
	kthSmallestInner(root)
	return value
}
