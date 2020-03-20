package graph

import (
	"testing"
	"fmt"
)

func TestGraph(t *testing.T) {
	graph := NewGraph(8)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 7)

	path := graph.BFS(0, 7)
	fmt.Println("BFS(0-7):")
	fmt.Println(path)

	path = graph.BFS(1,5)
	fmt.Println("BFS(1-5):")
	fmt.Println(path)

	path = graph.BFS(2,6)
	fmt.Println("BFS(2-6):")
	fmt.Println(path)

	path = graph.DFS(0, 7)
	fmt.Println("DFS(0-7):")
	fmt.Println(path)

	path = graph.DFS(1,5)
	fmt.Println("DFS(1-5):")
	fmt.Println(path)

	path = graph.DFS(2,6)
	fmt.Println("DFS(2-6):")
	fmt.Println(path)
}