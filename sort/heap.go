package sort

func Heap(b []byte) []byte {
	h := newHeap(b)
	h.sort(len(b))

	return b
}

type heap struct {
	b []byte
}

func newHeap(b []byte) *heap {
	heap := &heap{
		b: b,
	}

	return heap
}

func (h *heap) leftchildIndex(index int) int {
	return index<<1 + 1
}

func (h *heap) rightchildIndex(index int) int {
	return index<<1 + 2
}

func (h *heap) swap(first, second int) {
	h.b[first], h.b[second] = h.b[second], h.b[first]
}

func (h *heap) leaf(index int, size int) bool {
	return index >= (size>>1) && index <= size
}

func (h *heap) downHeapify(current int, size int) {
	if h.leaf(current, size) {
		return
	}

	biggest := current
	leftChildIndex := h.leftchildIndex(current)
	rightRightIndex := h.rightchildIndex(current)

	if leftChildIndex < size && h.b[leftChildIndex] > h.b[biggest] {
		biggest = leftChildIndex
	}

	if rightRightIndex < size && h.b[rightRightIndex] > h.b[biggest] {
		biggest = rightRightIndex
	}

	if biggest != current {
		h.swap(current, biggest)
		h.downHeapify(biggest, size)
	}
}

func (h *heap) buildHeap(size int) {
	for index := ((size >> 1) - 1); index >= 0; index-- {
		h.downHeapify(index, size)
	}
}

func (h *heap) sort(size int) {
	h.buildHeap(size)

	for i := size - 1; i > 0; i-- {
		// Move current root to end
		h.swap(0, i)
		h.downHeapify(0, i)
	}
}
