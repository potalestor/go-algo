package graph

import (
	"fmt"
)

type ID int

type Vertex struct {
	ID ID
}

func (v Vertex) String() string {
	return fmt.Sprintf("%v", v.ID)
}

type Edge struct {
	From   Vertex
	To     Vertex
	Weight float32
}

func (e Edge) String() string {
	if e.Weight == 0.0 {
		return fmt.Sprintf("%v-%v", e.From, e.To)
	}

	return fmt.Sprintf("%v-%v (%.1f)", e.From, e.To, e.Weight)
}

type Graph map[Vertex][]Edge

func NewGraph() Graph {
	return Graph{}
}

func (g Graph) AddVertex(v Vertex) {
	if _, ok := g[v]; ok {
		return
	}

	g[v] = []Edge{}
}

func (g Graph) AddEdge(u, v Vertex, w ...float32) {
	if len(w) == 0 {
		w = append(w, 0.0)
	}

	edge := Edge{From: u, To: v, Weight: w[0]}

	_, ok := g[u]
	if ok {
		g[u] = append(g[u], edge)
	} else {
		g[u] = []Edge{edge}
	}
}

func (g Graph) GetEdges() []Edge {
	result := []Edge{}

	for vertex, edges := range g {
		if len(edges) == 0 {
			result = append(result, Edge{From: vertex})

			continue
		}

		result = append(result, edges...)
	}

	return result
}
