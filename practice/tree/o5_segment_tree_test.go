package tree

import (
	"fmt"
	"testing"
)

type SegmentTree struct {
	tree []int
	data []int
	f    func(int, int) int
}

func NewSegmentTree(arr []int, f func(int, int) int) *SegmentTree {
	segmentTree := &SegmentTree{
		tree: make([]int, 4*len(arr)), // 极限高度差为1 所以系数为4
		data: make([]int, len(arr)),
		f:    f,
	}
	copy(segmentTree.data, arr)
	segmentTree.buildSegmentTree(0, 0, len(arr)-1)
	return segmentTree
}

func (segmentTree *SegmentTree) buildSegmentTree(treeIndex, left, right int) {
	if left == right {
		segmentTree.tree[treeIndex] = segmentTree.data[left]
		return
	}
	leftChildIndex := treeIndex<<1 + 1
	rightChildIndex := leftChildIndex + 1
	mid := ((right - left) >> 1) + left

	segmentTree.buildSegmentTree(leftChildIndex, left, mid)
	segmentTree.buildSegmentTree(rightChildIndex, mid+1, right)
	segmentTree.tree[treeIndex] = segmentTree.f(segmentTree.tree[leftChildIndex], segmentTree.tree[rightChildIndex])
}

func (segmentTree *SegmentTree) query(queryL, queryR int) int {
	if queryL < 0 || queryL >= len(segmentTree.data) ||
		queryR < 0 || queryR >= len(segmentTree.data) || queryL > queryR {
		panic("Index is illegal.")
	}

	var q func(int, int, int, int, int) int
	q = func(treeIndex, left, right, qL, qR int) int {
		if left == qL && right == qR {
			return segmentTree.tree[treeIndex]
		}
		mid := (right-left)>>1 + left
		leftChildIndex := treeIndex<<1 + 1
		rightChildIndex := leftChildIndex + 1
		if qL >= mid+1 {
			return q(rightChildIndex, mid+1, right, qL, qR)
		} else if qR <= mid {
			return q(leftChildIndex, left, mid, qL, qR)
		}
		leftRet := q(leftChildIndex, left, mid, qL, mid)
		rightRet := q(rightChildIndex, mid+1, right, mid+1, qR)
		return segmentTree.f(leftRet, rightRet)
	}
	return q(0, 0, len(segmentTree.data)-1, queryL, queryR)
}

func TestSegmentTree(t *testing.T) {
	nums := []int{1, 2, -7, -10, 23, 26, 11}
	tree := NewSegmentTree(nums, func(i int, i2 int) int {
		if i < i2 {
			return i
		} else {
			return i2
		}
		//return i + i2
	})
	//fmt.Println(tree.tree)
	fmt.Println(tree.query(0, 5))

}
