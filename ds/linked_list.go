// File: linked_list.go
// Language: Go

// This file contains the implementation of a linked list data structure.
// The List struct represents the linked list and the Node struct represents each node in the list.
// The List struct has methods to add a new element to the list, delete an element from the list,
// print the elements in the list, and get the length of the list.

package ds

import "fmt"

// List represents a linked list.
type List[T comparable] struct {
	head   *Node[T]
	length int
}

// Node represents a node in the linked list.
type Node[T comparable] struct {
	value T
	next  *Node[T]
}

// newNode creates a new node with the given value.
func newNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		value: value,
	}
}

// Add adds a new element to the end of the list.
func (l *List[T]) Add(value T) {
	node := newNode(value)
	if l.head == nil {
		l.head = node
	} else {
		currNode := l.head
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = node
	}
	l.length++
}

// Delete deletes the first occurrence of the given value from the list.
func (l *List[T]) Delete(value T) {
	if l.head == nil {
		fmt.Println("Empty List")
		return
	}
	if l.head.value == value {
		l.head = l.head.next
		l.length--
	} else {
		prev, curr := l.head, l.head.next
		for curr != nil {
			if curr.value == value {
				prev.next = curr.next
				l.length--
				return
			}
			prev, curr = curr, curr.next
		}
	}
	fmt.Println("Element not found in list")
}

// Print prints the elements in the list.
func (l *List[T]) Print() {
	if l.head == nil {
		fmt.Println("Empty List")
		return
	}
	currNode := l.head
	for currNode != nil {
		fmt.Print(currNode.value, "->")
		currNode = currNode.next
	}
	fmt.Println()
}

// Length returns the length of the list.
func (l *List[T]) Length() int {
	return l.length
}
