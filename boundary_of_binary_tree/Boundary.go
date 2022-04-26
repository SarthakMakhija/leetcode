package boundary_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func boundaryOfBinaryTree(root *TreeNode) []int {
	var leftBoundary []int
	var leftLeaves []int
	var rightLeaves []int
	var rightBoundary []int

	var leftBoundaryAndLeaves func(node *TreeNode, isInLeftBoundary bool)
	var rightBoundaryAndLeaves func(node *TreeNode, isInRightBoundary bool)

	leftBoundaryAndLeaves = func(node *TreeNode, isInLeftBoundary bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			leftLeaves = append(leftLeaves, node.Val)
		} else if isInLeftBoundary {
			leftBoundary = append(leftBoundary, node.Val)
		}
		leftBoundaryAndLeaves(node.Left, isInLeftBoundary)
		if isInLeftBoundary {
			if node.Left == nil {
				leftBoundaryAndLeaves(node.Right, true)
			} else {
				leftBoundaryAndLeaves(node.Right, false)
			}
		} else {
			leftBoundaryAndLeaves(node.Right, false)
		}
	}

	rightBoundaryAndLeaves = func(node *TreeNode, isInRightBoundary bool) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			rightLeaves = append(rightLeaves, node.Val)
		} else if isInRightBoundary {
			rightBoundary = append(rightBoundary, node.Val)
		}
		rightBoundaryAndLeaves(node.Right, isInRightBoundary)
		if isInRightBoundary {
			if node.Right == nil {
				rightBoundaryAndLeaves(node.Left, true)
			} else {
				rightBoundaryAndLeaves(node.Left, false)
			}
		} else {
			rightBoundaryAndLeaves(node.Left, false)
		}
	}

	reverse := func(numbers []int) []int {
		for i := 0; i < len(numbers)/2; i++ {
			j := len(numbers) - i - 1
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
		return numbers
	}

	leftBoundaryAndLeaves(root.Left, true)
	rightBoundaryAndLeaves(root.Right, true)
	leaves := append(leftLeaves, reverse(rightLeaves)...)

	return append(append([]int{root.Val}, append(leftBoundary, leaves...)...), reverse(rightBoundary)...)
}
