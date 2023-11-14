package ds

import (
	"fmt"
	"testing"
)

func TestList_Add(t *testing.T) {
	list := List[int]{}

	// Positive test case
	list.Add(1)
	if list.Length() != 1 {
		t.Errorf("Expected length to be 1, but got %d", list.Length())
	}

	// Negative test case
	list.Add(2)
	if list.Length() != 2 {
		t.Errorf("Expected length to be 2, but got %d", list.Length())
	}
}

func TestList_Delete(t *testing.T) {
	list := List[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Positive test case
	list.Delete(2)
	if list.Length() != 2 {
		t.Errorf("Expected length to be 2, but got %d", list.Length())
	}

	// Negative test case
	list.Delete(4)
	if list.Length() != 2 {
		t.Errorf("Expected length to be 2, but got %d", list.Length())
	}
}

func TestList_Print(t *testing.T) {
	list := List[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Positive test case
	fmt.Println("Expected output: 1->2->3->")
	list.Print()

	// Negative test case
	emptyList := List[int]{}
	fmt.Println("Expected output: Empty List")
	emptyList.Print()
}

func TestList_Length(t *testing.T) {
	list := List[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Positive test case
	if list.Length() != 3 {
		t.Errorf("Expected length to be 3, but got %d", list.Length())
	}

	// Negative test case
	emptyList := List[int]{}
	if emptyList.Length() != 0 {
		t.Errorf("Expected length to be 0, but got %d", emptyList.Length())
	}
}
