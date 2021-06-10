package graph

//https://www.imooc.com/video/12025
import (
	"fmt"
	"testing"
)

// 顶点
type Vertex struct {
	Data    interface{} // 顶点数据
	Visited bool
}

// 边
type Edge struct {
	VertexIndexA, VertexIndexB int
	Weight                     int
	Selected                   bool
}

// 图
type Graph struct {
	Capacity    int
	Size        int
	VertexArray []*Vertex
	EdgeArray   []*Edge
	Matrix      [][]int
}

// ==============================================、
func NewGraph(capacity int) *Graph {
	// 初始化矩阵
	matrix := make([][]int, capacity)
	for i := 0; i < capacity; i++ {
		matrix[i] = make([]int, capacity)
	}
	return &Graph{
		Capacity:    capacity,
		Size:        0,
		VertexArray: make([]*Vertex, capacity),
		EdgeArray:   make([]*Edge, 0),
		Matrix:      matrix,
	}
}

// 添加一个顶点
func (graph *Graph) AddVertex(v *Vertex) bool {
	if v == nil {
		return false
	}
	graph.VertexArray[graph.Size] = &*v
	graph.Size++
	return true
}

func (graph *Graph) Reset() {
	for i := 0; i < len(graph.VertexArray); i++ {
		graph.VertexArray[i].Visited = false
	}
}

func (graph *Graph) SetValue2Matrix4DirectedGraph(row, col, val int) bool {
	if row < 0 || col < 0 || row >= graph.Capacity || col >= graph.Capacity {
		return false
	}
	graph.Matrix[row][col] = val
	return true
}

func (graph *Graph) SetValue2Matrix4UnDirectedGraph(row, col, val int) bool {
	return graph.SetValue2Matrix4DirectedGraph(row, col, val) && graph.SetValue2Matrix4DirectedGraph(col, row, val)
}

func (graph *Graph) GetValueFromMatrix(row, col int) (bool, int) {
	if row < 0 || col < 0 || row >= graph.Capacity || col >= graph.Capacity {
		return false, 0
	}
	return true, graph.Matrix[row][col]
}

func (graph *Graph) PrintMatrix() {
	for i := 0; i < graph.Capacity; i++ {
		for j := 0; j < graph.Capacity; j++ {
			fmt.Printf("%d  ", graph.Matrix[i][j])
		}
		fmt.Println()
	}
}

//  获取最小边的索引
func (graph *Graph) GetMinEdgeIndex(edgeArr []*Edge) int {
	var minWeight, edgeIndex = 0, -1
	var i = 0

	// 找到第一条没有被访问过的边
	for ; i < len(edgeArr); i++ {
		if !edgeArr[i].Selected {
			minWeight = edgeArr[i].Weight
			edgeIndex = i
			break
		}
	}
	if edgeIndex == -1 {
		return edgeIndex
	}
	for ; i < len(edgeArr); i++ {
		if !edgeArr[i].Selected && edgeArr[i].Weight < minWeight {
			minWeight = edgeArr[i].Weight
			edgeIndex = i
		}
	}
	return edgeIndex
}

func isInside(vertexArr []int, index int) bool {
	for i := 0; i < len(vertexArr); i++ {
		if vertexArr[i] == index {
			return true
		}
	}
	return false
}

func mergeVertexSet(vertexSet [][]int, i, j int) {
	vertexSet[j] = append(vertexSet[j], vertexSet[i]...)
	vertexSet = append(vertexSet[:i], vertexSet[i+1:]...)
}

func (graph *Graph) DFS(vertexIndex int) {
	defer func() {
		graph.Reset()
	}()
	var process func(int)
	process = func(index int) {
		fmt.Printf("%s ", graph.VertexArray[index].Data.(string))
		graph.VertexArray[index].Visited = true
		// 找到与当前节点想链接的另一个顶点，并将状态向下传递
		for i := 0; i < graph.Capacity; i++ {
			if ok, val := graph.GetValueFromMatrix(index, i); ok && val != 0 && !graph.VertexArray[i].Visited {
				process(i)
			}
		}
	}
	process(vertexIndex)
}

func (graph *Graph) BFS(vertexIndex int) {
	defer func() {
		graph.Reset()
	}()
	fmt.Printf("%s ", graph.VertexArray[vertexIndex].Data.(string))
	graph.VertexArray[vertexIndex].Visited = true
	var process func([]int)
	process = func(prevLineIndexes []int) {
		currLineIndexes := make([]int, 0)
		for i := 0; i < len(prevLineIndexes); i++ {
			for j := 0; j < graph.Capacity; j++ {
				// 输出和prevLineIndexes链接的所有顶点
				if ok, val := graph.GetValueFromMatrix(prevLineIndexes[i], j); ok && val != 0 && !graph.VertexArray[j].Visited {
					fmt.Printf("%s ", graph.VertexArray[j].Data.(string))
					graph.VertexArray[j].Visited = true
					currLineIndexes = append(currLineIndexes, j)
				}
			}
		}
		if len(currLineIndexes) > 0 {
			process(currLineIndexes)
		}
	}
	process([]int{vertexIndex})
}

// 普利姆最小生成树
func (graph *Graph) PrimTree(vertexIndex int) {
	defer func() {
		graph.EdgeArray = make([]*Edge, 0)
		graph.Reset()
	}()
	// 已选点集合
	vertexIndexes := make([]int, 0)
	// 待选边集合
	edgeArr := make([]*Edge, 0)

	vertexIndexes = append(vertexIndexes, vertexIndex)
	graph.VertexArray[vertexIndex].Visited = true

	for len(graph.EdgeArray) < graph.Capacity-1 {
		lastVertexIndex := vertexIndexes[len(vertexIndexes)-1]
		// 和lastVertexIndex相关的边全部放到待选边集合中
		for i := 0; i < graph.Capacity; i++ {
			if ok, val := graph.GetValueFromMatrix(lastVertexIndex, i); ok && val != 0 && !graph.VertexArray[i].Visited {
				edgeArr = append(edgeArr, &Edge{
					VertexIndexA: lastVertexIndex,
					VertexIndexB: i,
					Weight:       val,
					Selected:     false,
				})
			}
		}

		minEdgeIndex := graph.GetMinEdgeIndex(edgeArr)
		edgeArr[minEdgeIndex].Selected = true

		fmt.Printf("%s --[%d]-- %s \n", graph.VertexArray[edgeArr[minEdgeIndex].VertexIndexA].Data.(string), edgeArr[minEdgeIndex].Weight, graph.VertexArray[edgeArr[minEdgeIndex].VertexIndexB].Data.(string))

		// 将待选边放到已选边中
		graph.EdgeArray = append(graph.EdgeArray, edgeArr[minEdgeIndex])

		nextVertexIndex := edgeArr[minEdgeIndex].VertexIndexB
		vertexIndexes = append(vertexIndexes, nextVertexIndex)
		graph.VertexArray[nextVertexIndex].Visited = true
	}
}

// 克鲁斯卡尔最小生成树
func (graph *Graph) KruskalTree() {
	vertexSet := make([][]int, 0)
	defer func() {
		graph.EdgeArray = make([]*Edge, 0)
		graph.Reset()
	}()

	// step1 取出所有边
	edgeArr := make([]*Edge, 0)
	// 取出矩阵主对角线上半部分(不含对角线)
	for r := 0; r < graph.Capacity; r++ {
		for c := r + 1; c < graph.Capacity; c++ {
			if ok, val := graph.GetValueFromMatrix(r, c); ok && val != 0 {
				// 构建边
				edgeArr = append(edgeArr, &Edge{
					VertexIndexA: r,
					VertexIndexB: c,
					Weight:       val,
					Selected:     false,
				})
			}
		}
	}
	// step2 从所有边中取出组成最小生成树的边
	// 2-1 找到结束条件
	for len(graph.EdgeArray) < graph.Capacity-1 {
		// 2-2 从集合中找到最小边
		minEdgeIndex := graph.GetMinEdgeIndex(edgeArr)
		edgeArr[minEdgeIndex].Selected = true
		// 2-3 找到最小边的连接点
		vertexIndexA, vertexIndexB := edgeArr[minEdgeIndex].VertexIndexA, edgeArr[minEdgeIndex].VertexIndexB
		vertexAIsInside, vertexBIsInside := false, false
		vertexAInsideIndex, vertexBInsideIndex := -1, -1

		// 2-4 找出点所在的点集合
		for i := 0; i < len(vertexSet); i++ {
			vertexAIsInside = isInside(vertexSet[i], vertexIndexA)
			vertexBIsInside = isInside(vertexSet[i], vertexIndexB)
			if vertexAIsInside {
				vertexAInsideIndex = i
			}
			if vertexBIsInside {
				vertexBInsideIndex = i
			}
		}
		// 2-5 根据点所在集合的不同做出不通处理
		if vertexAInsideIndex == -1 && vertexBInsideIndex == -1 {
			slice := []int{vertexIndexA, vertexIndexB}
			vertexSet = append(vertexSet, slice)
		} else if vertexAInsideIndex != -1 && vertexBInsideIndex != -1 {
			if vertexAInsideIndex != vertexBInsideIndex {
				mergeVertexSet(vertexSet, vertexAInsideIndex, vertexAInsideIndex)
			}
		} else if vertexAInsideIndex != -1 || vertexBInsideIndex != -1 {
			if vertexAInsideIndex != -1 {
				vertexSet[vertexAInsideIndex] = append(vertexSet[vertexAInsideIndex], vertexIndexB)
			} else {
				vertexSet[vertexBInsideIndex] = append(vertexSet[vertexBInsideIndex], vertexIndexA)
			}
		}
		graph.EdgeArray = append(graph.EdgeArray, edgeArr[minEdgeIndex])
		fmt.Printf("%s --[%d]-- %s \n", graph.VertexArray[edgeArr[minEdgeIndex].VertexIndexA].Data.(string), edgeArr[minEdgeIndex].Weight, graph.VertexArray[edgeArr[minEdgeIndex].VertexIndexB].Data.(string))
	}

}

func TestGraph(t *testing.T) {
	graph := NewGraph(6)
	A := &Vertex{Data: "A"}
	B := &Vertex{Data: "B"}
	C := &Vertex{Data: "C"}
	D := &Vertex{Data: "D"}
	E := &Vertex{Data: "E"}
	F := &Vertex{Data: "F"}

	graph.AddVertex(A)
	graph.AddVertex(B)
	graph.AddVertex(C)
	graph.AddVertex(D)
	graph.AddVertex(E)
	graph.AddVertex(F)

	graph.SetValue2Matrix4UnDirectedGraph(0, 1, 6)
	graph.SetValue2Matrix4UnDirectedGraph(0, 4, 5)
	graph.SetValue2Matrix4UnDirectedGraph(0, 5, 1)
	graph.SetValue2Matrix4UnDirectedGraph(1, 2, 3)
	graph.SetValue2Matrix4UnDirectedGraph(1, 5, 2)
	graph.SetValue2Matrix4UnDirectedGraph(2, 5, 8)
	graph.SetValue2Matrix4UnDirectedGraph(2, 3, 7)
	graph.SetValue2Matrix4UnDirectedGraph(3, 5, 4)
	graph.SetValue2Matrix4UnDirectedGraph(3, 4, 2)
	graph.SetValue2Matrix4UnDirectedGraph(4, 5, 9)

	graph.PrintMatrix()
	println("****** DFS ******")
	graph.DFS(0)
	println("\n****** BFS ******")
	graph.BFS(0)
	println("\n****** Prim ******")
	graph.PrimTree(0)
	println("\n****** Kruskal ******")
	graph.KruskalTree()

}
