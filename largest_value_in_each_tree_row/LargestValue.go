package largest_value_in_each_tree_row

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var largestValuesInner func(node *TreeNode, depth int)
	largestValueByDepth := make(map[int]int)

	largestValuesInner = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		largestValuesInner(node.Left, depth+1)
		largestValuesInner(node.Right, depth+1)

		largestValue, ok := largestValueByDepth[depth]
		if !ok {
			largestValueByDepth[depth] = node.Val
		} else if largestValue <= node.Val {
			largestValueByDepth[depth] = node.Val
		}
		return
	}
	largestValuesInner(root, 1)
	return toSlice(largestValueByDepth)
}

func toSlice(largestValueByDepth map[int]int) []int {
	depths := sortByKey(largestValueByDepth)

	var largestValues []int
	for _, depth := range depths {
		largestValues = append(largestValues, largestValueByDepth[depth])
	}
	return largestValues
}

func sortByKey(largestValueByDepth map[int]int) []int {
	depths := make([]int, 0, len(largestValueByDepth))
	for depth := range largestValueByDepth {
		depths = append(depths, depth)
	}
	sort.Ints(depths)
	return depths
}
