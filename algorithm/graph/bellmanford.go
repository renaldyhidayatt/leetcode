package graph

import (
	"errors"
	"math"
)

func (g *Graph) BellmanFord(start, end int) (isReachable bool, distance int, err error) {
	INF := math.Inf(1)
	distances := make([]float64, g.vertices)

	for i := 0; i < g.vertices; i++ {
		distances[i] = INF
	}
	distances[start] = 0

	for n := 0; n < g.vertices; n++ {
		for u, adjacents := range g.edges {
			for v, weightUV := range adjacents {
				if newDistance := distances[u] + float64(weightUV); distances[v] > newDistance {
					distances[v] = newDistance
				}
			}
		}
	}

	for u, adjacents := range g.edges {
		for v, weightUV := range adjacents {
			if newDistance := distances[u] + float64(weightUV); distances[v] > newDistance {
				return false, -1, errors.New("negative weight cycle present")
			}
		}
	}

	return distances[end] != INF, int(distances[end]), nil
}
