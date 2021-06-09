package graph

import (
	"gobyexample/practice/datastructure"
	"testing"
)

func TestGraph(t *testing.T) {

	graph := datastructure.NewGraph(6)
	A := &datastructure.Vertex{Data: "A"}
	B := &datastructure.Vertex{Data: "B"}
	C := &datastructure.Vertex{Data: "C"}
	D := &datastructure.Vertex{Data: "D"}
	E := &datastructure.Vertex{Data: "E"}
	F := &datastructure.Vertex{Data: "F"}

	graph.AddNode(A)
	graph.AddNode(B)
	graph.AddNode(C)
	graph.AddNode(D)
	graph.AddNode(E)
	graph.AddNode(F)

	graph.SetValueToMatrixForUndirectedGraph(0, 1, 6)
	graph.SetValueToMatrixForUndirectedGraph(0, 4, 5)
	graph.SetValueToMatrixForUndirectedGraph(0, 5, 1)
	graph.SetValueToMatrixForUndirectedGraph(1, 2, 3)
	graph.SetValueToMatrixForUndirectedGraph(1, 5, 2)
	graph.SetValueToMatrixForUndirectedGraph(2, 5, 8)
	graph.SetValueToMatrixForUndirectedGraph(2, 3, 7)
	graph.SetValueToMatrixForUndirectedGraph(3, 5, 4)
	graph.SetValueToMatrixForUndirectedGraph(3, 4, 2)
	graph.SetValueToMatrixForUndirectedGraph(4, 5, 9)

	graph.PrintMatrix()
	println("****** DFS ******")
	graph.DFS(0)
	println("\n****** BFS ******")
	graph.BFS(0)
	println("\n****** Prim ******")
	graph.PrimTree(0)

}
