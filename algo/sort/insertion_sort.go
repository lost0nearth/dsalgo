package sort

import "fmt"

func InsertionSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
		fmt.Printf("Intermediate Array %v\n", a)
	}
}
