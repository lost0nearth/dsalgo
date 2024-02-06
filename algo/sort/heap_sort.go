package sort

import "fmt"

func HeapSort(A []int) {
	n := len(A)

	i := n/2 - 1

	for i >= 0 {
		heapify(A, n, i)
		i--
	}

	fmt.Printf("Heapified Array %v\n", A)

	for i = n - 1; i > 0; i-- {
		A[0], A[i] = A[i], A[0]
		heapify(A, i, 0)
		fmt.Printf("Intermediate Array %v\n", A)
	}
}

func heapify(A []int, n, i int) {
	l := 2*i + 1
	r := 2*i + 2

	largest := i
	if l < n && A[l] > A[largest] {
		largest = l
	}

	if r < n && A[r] > A[largest] {
		largest = r
	}

	if largest != i {
		A[largest], A[i] = A[i], A[largest]
		heapify(A, n, largest)
	}
}
