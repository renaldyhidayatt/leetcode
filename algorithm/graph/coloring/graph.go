package coloring

import "errors"

type Color int

type Graph struct {
	Vertices int
	Edges    map[int]map[int]struct{}
}

func (g *Graph) AddVertex(v int) {
	if g.Edges == nil {
		g.Edges = map[int]map[int]struct{}{}
	}

	if _, ok := g.Edges[v]; !ok {
		g.Vertices++
		g.Edges[v] = make(map[int]struct{})
	}

}

func (g *Graph) AddEdge(one, two int) {
	g.AddVertex(one)
	g.AddVertex(two)

	g.Edges[one][two] = struct{}{}
	g.Edges[two][one] = struct{}{}
}

func (g *Graph) ValidateColorsOfVertex(colors map[int]Color) error {
	if g.Vertices != len(colors) {
		return errors.New("coloring: not all vertices of graph are colored")
	}

	for vertex, neighbours := range g.Edges {
		for nb := range neighbours {
			if colors[vertex] == colors[nb] {
				return errors.New("coloring: vertex and its neighbour have the same color")
			}
		}
	}

	return nil
}
