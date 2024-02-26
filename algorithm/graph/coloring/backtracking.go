package coloring

func (g *Graph) ColorUsingBacktracking() (map[int]Color, int) {
	vertexColors := make(map[int]Color, g.Vertices)

	g.colorVertex(0, vertexColors)

	colorsUsed := 0

	for _, cr := range vertexColors {
		if colorsUsed < int(cr) {
			colorsUsed = int(cr)
		}
	}

	return vertexColors, colorsUsed
}

func (g *Graph) colorVertex(v int, color map[int]Color) bool {
	if len(color) == g.Vertices {
		return true
	}

	for cr := Color(1); cr <= Color(g.Vertices); cr++ {
		safe := true
		for nb := range g.Edges[v] {
			if crnb, ok := color[nb]; ok && crnb == cr {
				safe = false
				break
			}
		}
		if safe {
			color[v] = cr
			if g.colorVertex(v+1, color) {
				return true
			}
			delete(color, v)
		}
	}
	return false
}
