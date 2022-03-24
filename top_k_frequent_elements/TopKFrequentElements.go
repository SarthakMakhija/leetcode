package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	topKFrequent := newTopKFrequentNums(k, len(nums))
	topKFrequent.arrangeBasedOnFrequency(nums)

	return topKFrequent.allKFrequentNums()
}

type topKFrequentNums struct {
	positionCountPairs map[int]*positionCountPair
	heap               *binaryHeap
	k                  int
}

type positionCountPair struct {
	heapPosition int
	count        int
}

func newTopKFrequentNums(k int, totalNums int) *topKFrequentNums {
	return &topKFrequentNums{
		positionCountPairs: make(map[int]*positionCountPair),
		heap:               newBinaryHeap(totalNums + 1),
		k:                  k,
	}
}

func (topKFrequent *topKFrequentNums) arrangeBasedOnFrequency(nums []int) {
	for _, num := range nums {
		pair, ok := topKFrequent.positionCountPairs[num]
		if !ok {
			occurrence := 1
			currentHeapPosition := topKFrequent.heap.currentIndex
			topKFrequent.positionCountPairs[num] = &positionCountPair{heapPosition: currentHeapPosition, count: occurrence}
			topKFrequent.heap.addAt(currentHeapPosition, frequency{count: occurrence, num: num}, topKFrequent)
			topKFrequent.heap.currentIndex = topKFrequent.heap.currentIndex + 1
		} else {
			newOccurrence := 1 + pair.count
			topKFrequent.heap.addAt(pair.heapPosition, frequency{count: newOccurrence, num: num}, topKFrequent)
			topKFrequent.positionCountPairs[num] = &positionCountPair{heapPosition: pair.heapPosition, count: newOccurrence}
		}
	}
}

func (topKFrequent *topKFrequentNums) allKFrequentNums() []int {
	var nums []int
	for count := 1; count <= topKFrequent.k; count++ {
		nums = append(nums, topKFrequent.heap.delete())
	}
	return nums
}

type frequency struct {
	num   int
	count int
}

type binaryHeap struct {
	elements     []frequency
	currentIndex int
}

func newBinaryHeap(size int) *binaryHeap {
	return &binaryHeap{elements: make([]frequency, size), currentIndex: 1}
}

func (heap *binaryHeap) addAt(index int, freq frequency, topKFrequent *topKFrequentNums) {
	heap.elements[index] = freq
	currentIndex := index
	parentIndex := currentIndex / 2

	for currentIndex > 1 && (heap.elements[currentIndex].count > heap.elements[parentIndex].count) {

		backUpCurrent := heap.elements[currentIndex]
		heap.elements[currentIndex] = heap.elements[parentIndex]
		topKFrequent.positionCountPairs[heap.elements[parentIndex].num].heapPosition = currentIndex

		heap.elements[parentIndex] = backUpCurrent
		topKFrequent.positionCountPairs[backUpCurrent.num].heapPosition = parentIndex

		currentIndex = parentIndex
		parentIndex = currentIndex / 2
	}
}

func (heap *binaryHeap) delete() int {
	swap := func(aIndex int, bIndex int) {
		hold := heap.elements[aIndex]
		heap.elements[aIndex] = heap.elements[bIndex]
		heap.elements[bIndex] = hold
	}

	rootIndex := 1
	num := heap.elements[rootIndex].num

	heap.elements[rootIndex] = heap.elements[heap.currentIndex-1]
	heap.currentIndex = heap.currentIndex - 1

	for rootIndex <= (heap.currentIndex-1)/2 {
		leftChildCount := heap.elements[2*rootIndex].count
		rightChildCount := heap.elements[2*rootIndex+1].count

		if leftChildCount > rightChildCount {
			if heap.elements[rootIndex].count < leftChildCount {
				swap(rootIndex, 2*rootIndex)
				rootIndex = 2 * rootIndex
			} else {
				break
			}
		} else {
			if heap.elements[rootIndex].count < rightChildCount {
				swap(rootIndex, 2*rootIndex+1)
				rootIndex = 2*rootIndex + 1
			} else {
				break
			}
		}
	}
	return num
}
