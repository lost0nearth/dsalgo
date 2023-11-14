package main

import (
	"github.com/dsalgo/ds"
)

func main() {
	l := &ds.List[int]{}
	l.Add(5)
	// l.Add(1)
	// l.Add(4)
	// fmt.Println(l.Length())
	// l.Add(2)
	// fmt.Println(l.Length())
	// l.Add(3)
	l.Print()
	l.Delete(5)
	l.Print()
}
