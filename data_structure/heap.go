package data_structure

type Heap interface {
	BuildMin() ()
}

type heap struct {
	list []int
	size int
}

func BuildMinHeap(list []int) (heap, error) {
	heap := &heap{
		list: list,
		size: len(list),
	}

	for i := len(list) / 2; i >= 0; i-- {
		heap.MinHeapify(i)
	}

	return *heap, nil
}

// NOTE
// 
// root        | i = 0
// parent      | (i-1) / 2
// left child  | 2i + 1
// right child | 2i + 2
func (h *heap) MinHeapify(i int) {
	min := i
	left := 2 * i + 1
	right := 2 * i + 2

	if left < h.size && h.list[min] > h.list[left] {
		min = left
	}
	if right < h.size && h.list[min] > h.list[right] {
		min = right
	}
	if min != i {
		h.list[i], h.list[min] = h.list[min], h.list[i]
		h.MinHeapify(min)
	}
}

func (h *heap) Sort() {
	for i := len(h.list) - 1; i >= 0; i-- {
		h.list[0], h.list[i] = h.list[i], h.list[0]
		h.size--
		h.MinHeapify(0)
	}
}