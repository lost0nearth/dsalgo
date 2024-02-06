package sort

import "fmt"

func BubbleSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
		fmt.Printf("Intermediate Array %v\n", a)
	}
}
