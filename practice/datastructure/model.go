package datastructure

import "fmt"

/*
Model  -  FOR TEST
*/

//单向链表
type NodeList struct {
	Val  interface{}
	Next *NodeList
}

//双向链表
type LinkedListNode struct {
	Val        interface{}
	Prev, Next *LinkedListNode
}

// 二叉树
type TreeNode struct {
	Val         interface{}
	Left, Right *TreeNode
}

// 栈
type Stack struct {
	top interface{}
	arr []interface{}
}

// 队列
type Queue struct {
	arr []interface{}
}

// 顶点
type Vertex struct {
	Data    interface{}
	visited bool
}

// 边
type Edge struct {
	vertex1, vertex2 int
	weight           int
	selected         bool
}

// 图
type Graph struct {
	Capacity    int
	Size        int
	VertexArray []*Vertex
	Edge        []*Edge //最小生成树边集合
	Matrix      [][]int
}

// ***************************GRAPH相关实现**************************************

func NewGraph(capacity int) *Graph {
	matrix := make([][]int, capacity)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, capacity)
	}
	return &Graph{
		Capacity:    capacity,
		Size:        0,
		VertexArray: make([]*Vertex, capacity),
		Edge:        make([]*Edge, capacity-1),
		Matrix:      matrix,
	}
}
func (g *Graph) AddNode(v *Vertex) bool {
	if v == nil {
		return false
	}
	g.VertexArray[g.Size] = &*v
	g.Size++
	return true
}
func (g *Graph) Reset() {
	for i := 0; i < g.Size; i++ {
		g.VertexArray[i].visited = false
	}
}
func (g *Graph) SetValueToMatrixForDirectedGraph(row, col, val int) bool {
	if row < 0 || col < 0 || row >= g.Capacity || col >= g.Capacity {
		return false
	}
	g.Matrix[row][col] = val
	return true
}
func (g *Graph) SetValueToMatrixForUndirectedGraph(row, col, val int) bool {
	return g.SetValueToMatrixForDirectedGraph(row, col, val) && g.SetValueToMatrixForDirectedGraph(col, row, val)
}
func (g *Graph) getValueFromMatrix(row, col int) (bool, int) {
	if row < 0 || col < 0 || row >= g.Capacity || col >= g.Capacity {
		return false, -1
	}
	return true, g.Matrix[row][col]
}
func (g *Graph) PrintMatrix() {
	for i := 0; i < g.Capacity; i++ {
		for j := 0; j < g.Capacity; j++ {
			fmt.Printf("%d  ", g.Matrix[i][j])
		}
		fmt.Println()
	}
}
func (g *Graph) getMinEdgeIndex(edgeList []*Edge) int {
	var minWeight = 0
	var edgeIndex = -1
	var i = 0
	// 找到第一条没有被访问过的边
	for ; i < len(edgeList); i++ {
		if !edgeList[i].selected {
			minWeight = edgeList[i].weight
			edgeIndex = i
			break
		}
	}
	if edgeIndex == -1 {
		return edgeIndex
	}
	for ; i < len(edgeList); i++ {
		if !edgeList[i].selected && minWeight > edgeList[i].weight {
			minWeight = edgeList[i].weight
			edgeIndex = i
		}
	}
	return edgeIndex
}

func (g *Graph) DFS(nodeIndex int) {
	defer func() {
		g.Reset()
	}()
	var dfs func(int)
	dfs = func(idx int) {
		fmt.Printf("%s ", g.VertexArray[idx].Data.(string))
		g.VertexArray[idx].visited = true
		for i := 0; i < g.Capacity; i++ {
			// 找到关联的顶点 状态向下转移
			if ok, val := g.getValueFromMatrix(idx, i); ok && val != 0 && !g.VertexArray[i].visited {
				dfs(i)
			}
		}
	}
	dfs(nodeIndex)
}
func (g *Graph) BFS(nodeIndex int) {
	defer func() {
		g.Reset()
	}()
	fmt.Printf("%s ", g.VertexArray[nodeIndex].Data.(string))
	g.VertexArray[nodeIndex].visited = true

	var bfs func([]int)
	bfs = func(prevLineIndexes []int) {
		currLineIndexes := make([]int, 0)
		for i := 0; i < len(prevLineIndexes); i++ {
			for j := 0; j < g.Capacity; j++ {
				// 找到和prevLineIndexes[i]的直属关联顶点
				if ok, val := g.getValueFromMatrix(prevLineIndexes[i], j); ok && val != 0 && !g.VertexArray[j].visited {
					fmt.Printf("%s ", g.VertexArray[j].Data.(string))
					g.VertexArray[j].visited = true
					currLineIndexes = append(currLineIndexes, j)
				}
			}
		}
		if len(currLineIndexes) > 0 {
			bfs(currLineIndexes)
		}
	}
	bfs([]int{nodeIndex})
}

// 普利姆算法-最小生成树
func (g *Graph) PrimTree(index int) {
	defer func() {
		g.Reset()
	}()
	// 已选点集合
	vertexIndexes := make([]int, 0)
	// 已选边集合
	edgeList := make([]*Edge, 0)
	// 已选边数量
	edgeCount := 0
	// 初始化已选点
	vertexIndexes = append(vertexIndexes, index)

	g.VertexArray[index].visited = true

	// 当已选边和所有点满足 edgeCount == g.Capacity-1 时候退出
	for edgeCount < g.Capacity-1 {
		tmp := vertexIndexes[len(vertexIndexes)-1]
		// for循环结束后就将和顶点tmp相关的边全部放到备选边集合中
		for i := 0; i < g.Capacity; i++ {
			if ok, val := g.getValueFromMatrix(tmp, i); ok && val != 0 && !g.VertexArray[i].visited {
				// 放入备选边中
				edgeList = append(edgeList, &Edge{vertex1: tmp, vertex2: i, weight: val})
			}
		}
		// 从备选边集合中找到最小的边索引
		edgeIndex := g.getMinEdgeIndex(edgeList)
		// 标记为已选择
		edgeList[edgeIndex].selected = true

		fmt.Printf("%s --[%d]-- %s \n", g.VertexArray[edgeList[edgeIndex].vertex1].Data.(string), edgeList[edgeIndex].weight, g.VertexArray[edgeList[edgeIndex].vertex2].Data.(string))

		// 放到最小生成树的集合中
		g.Edge[edgeCount] = edgeList[edgeIndex]
		edgeCount++

		// 根据最小边找到连接的另一个点的索引
		nextNodeIndex := edgeList[edgeIndex].vertex2
		// 添加到已选的点中
		vertexIndexes = append(vertexIndexes, nextNodeIndex)
		// 标记为已选择
		g.VertexArray[nextNodeIndex].visited = true

	}
	fmt.Printf("\n顶点的选择顺序  ")
	for i := 0; i < len(vertexIndexes); i++ {
		fmt.Printf("%s ", g.VertexArray[vertexIndexes[i]].Data.(string))
	}
	fmt.Println()
}

// ***************************STACK相关实现**************************************
func (s *Stack) Push(val interface{}) {
	s.top = val
	s.arr = append(s.arr, val)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	ret := s.top
	s.arr = s.arr[:len(s.arr)-1]
	if len(s.arr) == 0 {
		s.top = nil
	} else {
		s.top = s.arr[len(s.arr)-1]
	}
	return ret
}

func (s *Stack) IsEmpty() bool {
	return s.arr == nil || len(s.arr) == 0
}

// *****************************QUEUE相关实现************************************
func (q *Queue) IsEmpty() bool {
	return q.arr == nil || len(q.arr) == 0
}

func (q *Queue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	ret := q.arr[0]
	q.arr = q.arr[1:]
	return ret
}

func (q *Queue) Add(val interface{}) {
	if q.arr == nil {
		q.arr = make([]interface{}, 0)
	}
	q.arr = append(q.arr, val)
}
