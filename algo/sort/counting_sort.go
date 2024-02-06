package sort

import "fmt"

func CountingSort(A []int) []int {
	k := getMaxValue(A)
	n := len(A)
	B := make([]int, n)
	C := make([]int, k+1)
	for j := range A {
		C[A[j]]++
	}

	fmt.Printf("Count Array %v\n", C)

	for i := 1; i <= k; i++ {
		C[i] = C[i] + C[i-1]
	}

	fmt.Printf("CumCount Array %v\n", C)

	// from n-1 to 1 is stable, but from 1 to n-1 is not stable. Easier to understand from context of radix sort
	// and radix sort needs stable
	for j := n - 1; j >= 0; j-- {
		fmt.Printf("j = %d A[j] = %d and C[A[j]] = %d\n", j, A[j], C[A[j]])

		C[A[j]]--
		B[C[A[j]]] = A[j]

		fmt.Printf("Result Array %v\n", B)

		fmt.Printf("CumCount Array %v\n", C)
	}
	return B
}
