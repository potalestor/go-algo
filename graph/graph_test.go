package graph

import (
	"fmt"
	"testing"
)

func TestGraph_String(t *testing.T) {
	g := NewGraph()

	g.Directed = true

	g.AddEdge(Vertex{1}, Vertex{2})
	g.AddEdge(Vertex{2}, Vertex{3})
	g.AddEdge(Vertex{2}, Vertex{4})
	g.AddEdge(Vertex{1}, Vertex{4})

	g.AddVertex(Vertex{1})

	fmt.Println(g)
}
