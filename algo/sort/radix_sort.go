package sort

import "fmt"

func RadixSort(A []int) []int {
	maxValue := getMaxValue(A)

	for exp := 1; maxValue/exp > 0; exp *= 10 {
		A = countingSortByDigit(A, exp)
		fmt.Printf("Intermediate Array %v\n", A)
	}

	return A
}

// getMaxValue returns the maximum value in the input array A.
func getMaxValue(A []int) int {
	max := A[0]
	for _, num := range A {
		if num > max {
			max = num
		}
	}
	return max
}

// countingSortByDigit performs counting sort on the input array A based on the digit at the given exponent.
func countingSortByDigit(A []int, exp int) []int {
	n := len(A)
	B := make([]int, n)
	C := make([]int, 10) // We assume decimal digits (base 10)

	// Count occurrences of digits at the given exponent
	for i := 0; i < n; i++ {
		digit := (A[i] / exp) % 10
		C[digit]++
	}

	// Calculate cumulative counts
	for i := 1; i < 10; i++ {
		C[i] += C[i-1]
	}

	// Build the sorted array
	for i := n - 1; i >= 0; i-- {
		digit := (A[i] / exp) % 10
		C[digit]--
		B[C[digit]] = A[i]
	}

	return B
}
