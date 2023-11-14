package ds

import "fmt"

type List[T comparable] struct {
	head   *Node[T]
	length int
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func newNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		value: value,
	}
}

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

func (l *List[T]) Length() int {
	return l.length
}
