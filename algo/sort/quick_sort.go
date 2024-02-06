package sort

import "fmt"

func partition(A []int, p, r int) int {
	fmt.Printf("Intermediate Array %v ", A[p:r+1])
	pivot := A[r]
	i, j := p-1, p
	for j < r {
		if A[j] <= pivot {
			i++
			A[i], A[j] = A[j], A[i]
		}
		j++
	}
	i++
	A[i], A[r] = A[r], A[i]
	fmt.Printf("%v\n", A[p:r+1])
	return i
}

func QuickSort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		QuickSort(A, p, q-1)
		QuickSort(A, q+1, r)
	}
}
