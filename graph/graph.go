package graph

import (
	"fmt"
	"strings"
)

type ID int

type Edge struct {
	From   Vertex
	To     Vertex
	Weight float32
}

func (e *Edge) String() string {
	return fmt.Sprintf("%v-%v (%f)", e.From, e.To, e.Weight)
}

type Vertex struct {
	ID ID
}

func (v *Vertex) String() string {
	return fmt.Sprintf("%v", v.ID)
}

type Graph struct {
	Directed bool
	Edges    map[Vertex]map[Edge]struct{}
}

// NewGraph: Create graph.
func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[Vertex]map[Edge]struct{}),
	}
}

func (g *Graph) AddVertex(v Vertex) {
	if _, ok := g.Edges[v]; ok {
		return
	}

	g.Edges[v] = nil
}

func (g *Graph) AddEdge(u, v Vertex, w ...float32) {
	if len(w) == 0 {
		w = append(w, 0.0)
	}

	newEdge := Edge{From: u, To: v, Weight: w[0]}

	_, ok := g.Edges[u]
	if ok {
		g.Edges[u][newEdge] = struct{}{}
	} else {
		g.Edges[u] = map[Edge]struct{}{newEdge: {}}
	}

	if !g.Directed {
		newEdge = Edge{From: v, To: u, Weight: w[0]}

		_, ok := g.Edges[v]
		if ok {
			g.Edges[v][newEdge] = struct{}{}
		} else {
			g.Edges[v] = map[Edge]struct{}{newEdge: {}}
		}
	}
}

func (g *Graph) String() string {
	var b strings.Builder

	for vertex, edges := range g.Edges {
		if len(edges) == 0 {
			b.WriteString(vertex.String())
			b.WriteRune('\n')

			continue
		}

		for edge := range edges {
			b.WriteString(edge.String())
			b.WriteRune('\n')
		}
	}

	return b.String()
}
