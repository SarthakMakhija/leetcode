package path_sum2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var pathSumInner func(node *TreeNode, remainingSum int) [][]int

	pathSumInner = func(node *TreeNode, remainingSum int) [][]int {
		switch {
		case node == nil:
			return [][]int{}
		case node.Left == nil && node.Right == nil && remainingSum-node.Val == 0:
			return [][]int{{node.Val}}
		}
		merged := merge(pathSumInner(node.Left, remainingSum-node.Val), pathSumInner(node.Right, remainingSum-node.Val))
		if len(merged) != 0 {
			return addTo(merged, node.Val)
		}
		return merged
	}
	return pathSumInner(root, targetSum)
}

func merge(first [][]int, second [][]int) [][]int {
	merged := copy2D(first)
	for index := 0; index < len(second); index++ {
		merged = append(merged, second[index])
	}
	return merged
}

func copy2D(elements [][]int) [][]int {
	copied := make([][]int, len(elements))
	for index := range elements {
		copied[index] = make([]int, len(elements[index]))
		copy(copied[index], elements[index])
	}
	return copied
}

func addTo(elements [][]int, element int) [][]int {
	if len(elements) == 0 {
		return append(elements, []int{element})
	}
	copied := copy2D(elements)
	for index := 0; index < len(copied); index++ {
		replacement := make([]int, len(copied[index])+1)
		replacement[0] = element
		copy(replacement[1:], copied[index])
		copied[index] = replacement
	}
	return copied
}
