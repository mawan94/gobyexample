package unionfind

import "testing"

type UnionFind struct {
	parent []int
}

func NewUnionFind(arr []interface{}) *UnionFind {
	parent := make([]int, len(arr), len(arr))
	for i := 0; i < len(arr); i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
	}
}

func (uf *UnionFind) Union(indexes ...int) {
	for i := 1; i < len(indexes); i++ {
		uf.parent[uf.find(indexes[i])] = uf.find(indexes[0])
	}
}

func (uf *UnionFind) IsConnected(indexes ...int) bool {
	headRoot := uf.find(indexes[0])
	for i := 1; i < len(indexes); i++ {
		iRoot := uf.find(indexes[i])
		if headRoot != iRoot {
			return false
		}
	}
	return true
}

// 找到最上层的根
func (uf *UnionFind) find(index int) int {
	if index != uf.parent[index] {
		uf.parent[index] = uf.find(uf.parent[index])
	}
	return uf.parent[index]
}

func TestUnionFind(t *testing.T) {
	arr := []interface{}{"aa", "bb", "cc", "dd"}
	uf := NewUnionFind(arr)
	println(uf.IsConnected(0, 1))
	uf.Union(0, 1)
	println(uf.IsConnected(0, 1))
	println(uf.IsConnected(1, 2))
	uf.Union(1, 2)
	println(uf.IsConnected(1, 2))
	println(uf.IsConnected(1, 0))

}
