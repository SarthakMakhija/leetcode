package range_sum_query_mutable

//NumArray
//this is segment tree
type NumArray struct {
	size int
	tree []int
}

func Constructor(nums []int) NumArray {
	var build func(start, end, segmentTreeIndex int, nums []int)

	numArray := NumArray{size: len(nums), tree: make([]int, 4*len(nums))}
	build = func(start, end, segmentTreeIndex int, nums []int) {
		if start == end {
			numArray.tree[segmentTreeIndex] = nums[start]
			return
		}
		mid := (start + end) / 2
		build(start, mid, 2*segmentTreeIndex+1, nums)
		build(mid+1, end, 2*segmentTreeIndex+2, nums)
		numArray.tree[segmentTreeIndex] = numArray.tree[2*segmentTreeIndex+1] + numArray.tree[2*segmentTreeIndex+2]
	}
	build(0, numArray.size-1, 0, nums)
	return numArray
}

func (numArray *NumArray) Update(index int, val int) {
	var updateInner func(start, end, segmentTreeIndex int)
	updateInner = func(start, end, segmentTreeIndex int) {
		if start == end {
			numArray.tree[segmentTreeIndex] = val
			return
		}
		mid := (start + end) / 2
		if index <= mid {
			updateInner(start, mid, 2*segmentTreeIndex+1)
		} else {
			updateInner(mid+1, end, 2*segmentTreeIndex+2)
		}
		numArray.tree[segmentTreeIndex] = numArray.tree[2*segmentTreeIndex+1] + numArray.tree[2*segmentTreeIndex+2]
	}
	updateInner(0, numArray.size-1, 0)
}

func (numArray *NumArray) SumRange(left int, right int) int {
	var sumRangeInner func(start, end, segmentTreeIndex int) int
	sumRangeInner = func(start, end, segmentTreeIndex int) int {
		//no overlap
		if start > right || end < left {
			return 0
		}
		//complete overlap
		if start >= left && end <= right {
			return numArray.tree[segmentTreeIndex]
		}
		mid := (start + end) / 2
		return sumRangeInner(start, mid, 2*segmentTreeIndex+1) + sumRangeInner(mid+1, end, 2*segmentTreeIndex+2)
	}
	return sumRangeInner(0, numArray.size-1, 0)
}
