package ds

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Size() int {
	return len(h.array)
}

func (h *MaxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) leftChild(i int) int {
	return 2*i + 1
}

func (h *MaxHeap) rightChild(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) Heapify(i int) {
	l := h.leftChild(i)
	r := h.rightChild(i)

	largest := i

	if l < h.Size() && h.array[l] > h.array[largest] {
		largest = l
	}

	if r < h.Size() && h.array[r] > h.array[largest] {
		largest = r
	}

	if largest != i {
		h.array[i], h.array[largest] = h.array[largest], h.array[i]
		h.Heapify(largest)
	}
}

func (h *MaxHeap) Build(A []int) {
	n := len(A)
	h.array = append(h.array, A...)

	i := h.parent(n - 1)

	for i >= 0 {
		h.Heapify(i)
		i--
	}
}

func (h *MaxHeap) AddElement(element int) {
	h.array = append(h.array, element)
	index := h.Size() - 1

	for index > 0 && h.array[h.parent(index)] < h.array[index] {
		h.array[h.parent(index)], h.array[index] = h.array[index], h.array[h.parent(index)]
		index = h.parent(index)
	}
}

func (h MaxHeap) Print() {
	fmt.Printf("Heap Array %v\n", h.array)
}
