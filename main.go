package main

import (
	"fmt"

	"github.com/lost0nearth/dsalgo/algo/sort"
)

var a = []int{170, 45, 75, 90, 802, 24, 2, 66, 2021}

func TestInsertionSort() {
	fmt.Printf("Input Array %v\n", a)
	sort.InsertionSort(a)
	fmt.Printf("SortedArray %v\n", a)
}

func TestSelectionSort() {
	fmt.Printf("Input Array %v\n", a)
	sort.SelectionSort(a)
	fmt.Printf("SortedArray %v\n", a)
}

func TestMergeSort() {
	fmt.Printf("Input Array %v\n", a)
	sort.MergeSort(a, 0, len(a)-1)
	fmt.Printf("SortedArray %v\n", a)
}

func TestBubbleSort() {
	fmt.Printf("Input Array %v\n", a)
	sort.BubbleSort(a)
	fmt.Printf("SortedArray %v\n", a)
}

func TestHeapSort() {
	fmt.Printf("Input Array %v\n", a)
	sort.HeapSort(a)
	fmt.Printf("SortedArray %v\n", a)
}

func TestQuickSort() {
	fmt.Printf("Input Array %v\n", a)
	sort.QuickSort(a, 0, len(a)-1)
	fmt.Printf("SortedArray %v\n", a)
}

func TestCountingSort() {
	fmt.Printf("Input Array %v\n", a)
	b := sort.CountingSort(a)
	fmt.Printf("SortedArray %v\n", b)
}

func TestRadixSort() {
	fmt.Printf("Input Array %v\n", a)
	b := sort.RadixSort(a)
	fmt.Printf("SortedArray %v\n", b)
}

func main() {
	// TestBubbleSort()
	// TestSelectionSort()
	// TestInsertionSort()
	// TestMergeSort()
	// TestInsertionSort()
	// TestCountingSort()
	TestRadixSort()
}
