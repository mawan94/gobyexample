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
}
