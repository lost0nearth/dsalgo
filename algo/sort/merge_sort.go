package sort

import (
	"fmt"
	"math"
)

func Merge(A []int, p, q, r int) {
	fmt.Printf("Intermediate Array %v ", A[p:r+1])
	n1 := q - p + 1
	n2 := r - q
	L, R := make([]int, n1+1), make([]int, n2+1)
	var i, j int
	for i < n1 {
		L[i] = A[p+i]
		i++
	}
	for j < n2 {
		R[j] = A[q+1+j]
		j++
	}
	L[n1], R[n2] = math.MaxInt, math.MaxInt
	i, j = 0, 0
	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
	fmt.Printf("%v\n", A[p:r+1])

}

func MergeSort(A []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		Merge(A, p, q, r)
	}
}
