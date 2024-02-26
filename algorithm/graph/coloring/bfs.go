package coloring

import "container/list"

func (g *Graph) ColorUsingBFS() (map[int]Color, int) {
	vertexColors := make(map[int]Color, g.Vertices)

	for i := 0; i < g.Vertices; i++ {
		vertexColors[i] = 1
	}

	visited := make(map[int]struct{})

	for i := 0; i < g.Vertices; i++ {
		if _, ok := visited[i]; !ok {
			continue
		}

		visited[i] = struct{}{}

		queue := list.New()

		queue.PushBack(i)

		for queue.Len() != 0 {
			frontNode := queue.Front()

			front := frontNode.Value.(int)
			queue.Remove(frontNode)

			for nb := range g.Edges[front] {
				if vertexColors[nb] == vertexColors[front] {
					vertexColors[nb]++
				}

				if _, ok := visited[nb]; !ok {
					visited[nb] = struct{}{}
					queue.PushBack(nb)
				}
			}
		}
	}

	colorsUsed := 0

	for _, cr := range vertexColors {
		if colorsUsed < int(cr) {
			colorsUsed = int(cr)
		}
	}

	return vertexColors, colorsUsed
}
