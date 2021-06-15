package unionfind

import "testing"

/*
	作用: 查看两个元素是否有关联 （路径连通， 关系网络（是否可达））。
	并查集也可以看成是一个直接给出结果的【图】。

 	建立一种机制，对目标容器做一个模型映射，保证元素次序一一对应。
	映射容器保存上级索引（合并就是操作映射容器对应位置上的【值】，这个【值】也是其上级索引。）
*/
type UnionFind struct {
	parent []int
}

// 集合元素位置一一对应UnionFind.parent位置
// 初始化时候各自为营，每个元素的上级都是自己
func NewUnionFind(arr []interface{}) *UnionFind {
	parent := make([]int, len(arr), len(arr))
	for i := 0; i < len(arr); i++ {
		parent[i] = i
	}
	return &UnionFind{
		parent: parent,
	}
}

// 合并元素
func (uf *UnionFind) Union(indexes ...int) {
	// 其余参数的根节点全部指向第一个参数的父节点
	for i := 1; i < len(indexes); i++ {
		uf.parent[uf.find(indexes[i])] = uf.find(indexes[0])
	}
}

// 是否关联 (存在相同的root)
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

// 找到最上层的root索引
func (uf *UnionFind) find(index int) int {
	// 这里顺便做一下路径压缩
	if index != uf.parent[index] {
		uf.parent[index] = uf.find(uf.parent[index])
	}
	return uf.parent[index]
}

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind([]interface{}{"aa", "bb", "cc", "dd"})

	println(uf.IsConnected(0, 1))
	uf.Union(0, 1)
	println(uf.IsConnected(0, 1))

	println(uf.IsConnected(1, 2))
	uf.Union(1, 2)
	println(uf.IsConnected(2, 0, 1))
	println(uf.IsConnected(2, 0, 1, 3))
}
