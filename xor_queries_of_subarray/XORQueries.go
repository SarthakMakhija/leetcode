package xor_queries_of_subarray

func xorQueries(arr []int, queries [][]int) []int {
	segmentTree := ConstructSegmentTreeFrom(arr)
	var replies []int
	for _, query := range queries {
		replies = append(replies, segmentTree.xorQuery(query[0], query[1]))
	}
	return replies
}

type SegmentTree struct {
	tree         []int
	originalSize int
}

func ConstructSegmentTreeFrom(arr []int) SegmentTree {
	tree := SegmentTree{originalSize: len(arr), tree: make([]int, 4*len(arr))}

	var build func(start, end, segmentTreeIndex int)
	build = func(start, end, segmentTreeIndex int) {
		if start == end {
			tree.tree[segmentTreeIndex] = arr[start]
			return
		}
		mid := (start + end) / 2
		build(start, mid, 2*segmentTreeIndex+1)
		build(mid+1, end, 2*segmentTreeIndex+2)
		tree.tree[segmentTreeIndex] = tree.tree[2*segmentTreeIndex+1] ^ tree.tree[2*segmentTreeIndex+2]
	}
	build(0, tree.originalSize-1, 0)
	return tree
}

func (tree SegmentTree) xorQuery(left, right int) int {
	var xorQueryInner func(start, end, segmentTreeIndex int) int
	xorQueryInner = func(start, end, segmentTreeIndex int) int {
		//no overlap
		if start > right || end < left {
			return 0
		}
		//complete overlap
		if start >= left && end <= right {
			return tree.tree[segmentTreeIndex]
		}
		mid := (start + end) / 2
		return xorQueryInner(start, mid, 2*segmentTreeIndex+1) ^ xorQueryInner(mid+1, end, 2*segmentTreeIndex+2)
	}
	return xorQueryInner(0, tree.originalSize-1, 0)
}
