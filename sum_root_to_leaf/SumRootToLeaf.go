package sum_root_to_leaf

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeValue struct {
	val   int
	power int
}

func sumNumbers(root *TreeNode) int {
	var sumNumbersInner func(node *TreeNode) []TreeValue

	sumNumbersInner = func(node *TreeNode) []TreeValue {
		switch {
		case node == nil:
			return []TreeValue{}
		case node.Left == nil && node.Right == nil:
			return []TreeValue{{val: node.Val, power: 1}}
		}
		merged := merge(sumNumbersInner(node.Left), sumNumbersInner(node.Right))
		return multiply(node.Val, merged)
	}
	return addAllIn(sumNumbersInner(root))
}

func merge(one []TreeValue, other []TreeValue) []TreeValue {
	var merged []TreeValue
	partial := append(merged, one...)
	return append(partial, other...)
}

func multiply(value int, treeValues []TreeValue) []TreeValue {
	var multiplied []TreeValue
	for _, treeValue := range treeValues {
		multiplied = append(multiplied,
			TreeValue{
				power: treeValue.power + 1,
				val:   value*int(math.Pow10(treeValue.power)) + treeValue.val,
			})
	}
	return multiplied
}

func addAllIn(treeValues []TreeValue) int {
	sum := 0
	for _, treeValue := range treeValues {
		sum = sum + treeValue.val
	}
	return sum
}
