package binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	values       []int
	currentIndex int
}

func Constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	var iterate func(node *TreeNode) []int
	iterate = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}
		var values []int
		return append(append(append(values, iterate(node.Left)...), node.Val), iterate(node.Right)...)
	}
	iterator.values = iterate(root)
	iterator.currentIndex = 0
	return iterator
}

func (bstIterator *BSTIterator) Next() int {
	value := bstIterator.values[bstIterator.currentIndex]
	bstIterator.currentIndex = bstIterator.currentIndex + 1
	return value
}

func (bstIterator *BSTIterator) HasNext() bool {
	return bstIterator.currentIndex < len(bstIterator.values)
}
