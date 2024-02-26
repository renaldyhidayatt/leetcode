package coloring

import "sort"

func (g *Graph) ColorUsingGreedyApproach() (map[int]Color, int) {
	degreeOfVertex := make([]struct{ degree, vertex int }, 0, g.Vertices)
	for v, neighbours := range g.Edges {
		degreeOfVertex = append(degreeOfVertex,
			struct{ degree, vertex int }{
				vertex: v,
				degree: len(neighbours),
			},
		)
	}

	sort.Slice(degreeOfVertex, func(i, j int) bool {
		return degreeOfVertex[i].degree > degreeOfVertex[j].degree
	})

	vertexColors := make(map[int]Color, g.Vertices)
	for color := 1; color <= g.Vertices; color++ {
	vertexLoop:
		for _, val := range degreeOfVertex {
			if _, ok := vertexColors[val.vertex]; ok {
				continue vertexLoop
			}
			for ng := range g.Edges[val.vertex] {
				if vertexColors[ng] == Color(color) {
					continue vertexLoop
				}
			}
			vertexColors[val.vertex] = Color(color)
		}
		if len(vertexColors) == g.Vertices {
			return vertexColors, color
		}
	}
	return vertexColors, g.Vertices
}
