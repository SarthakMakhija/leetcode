package binary_tree_maximum_path_sum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt

	var maxPathSumInner func(node *TreeNode) int
	maxPathSumInner = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := maxPathSumInner(node.Left)
		rightSum := maxPathSumInner(node.Right)
		currentNodeAsRootSum := int(math.Max(
			float64(rightSum+node.Val),
			math.Max(
				float64(leftSum+node.Val),
				math.Max(float64(leftSum+node.Val+rightSum), float64(node.Val)),
			),
		))
		if currentNodeAsRootSum > maxSum {
			maxSum = currentNodeAsRootSum
		}
		return int(math.Max(float64(node.Val), math.Max(float64(node.Val+leftSum), float64(node.Val+rightSum))))
	}
	maxPathSumInner(root)
	return maxSum
}
