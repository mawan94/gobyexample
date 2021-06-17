package datastructure

import (
	"fmt"
	"testing"
)

func TestDataStruct(t *testing.T) {
	list := NewLinkedListBySlice([]interface{}{1, 2, 3, 4, 5})
	list.add(1, 888)

	list.print(func(e interface{}) {
		fmt.Printf("%d ", e)
	})

	println()
	list.add(5, 666)
	list.print(func(e interface{}) {
		fmt.Printf("%d ", e)
	})
	println()
	list.remove(6)
	list.remove(5)
	list.remove(0)
	list.print(func(e interface{}) {
		fmt.Printf("%d ", e)
	})
	list.Swap(0, 3)
	println()
	list.print(func(e interface{}) {
		fmt.Printf("%d ", e)
	})
	list.Swap(2, 1)
	println()
	list.print(func(e interface{}) {
		fmt.Printf("%d ", e)
	})

	// ********************* BST **************************
	bst := NewBinarySearchTree()
	var a Integer = 4
	var b Integer = 3
	var c Integer = 2
	var d Integer = 1

	bst.add(a)
	bst.add(b)
	bst.add(c)
	bst.add(d)

	bst.remove(d)
	bst.remove(a)

	bst.bfs(func(e interface{}) {
		fmt.Printf("%d ", e.(Integer))
	})
	bst.sortOrder(func(e interface{}) {
		fmt.Printf("%d ", e.(Integer))
	})
}

type Integer int

func (self Integer) compareTo(other Comparable) int {
	if self < other.(Integer) {
		return -1
	} else if self > other.(Integer) {
		return 1
	}
	return 0
}
