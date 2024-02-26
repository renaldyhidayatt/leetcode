package graph

import "math"

type WeightedGraph [][]float64

var Inf = math.Inf(1)

func FloydWarshall(graph WeightedGraph) WeightedGraph {
	if len(graph) == 0 || len(graph) != len(graph[0]) {
		return nil
	}

	for i := 0; i < len(graph); i++ {
		if len(graph[i]) != len(graph) {
			return nil
		}
	}

	numVertices := len(graph)

	result := make(WeightedGraph, numVertices)

	for i := 0; i < numVertices; i++ {
		result[i] = make([]float64, numVertices)

		for j := 0; j < numVertices; j++ {
			result[i][j] = graph[i][j]
		}
	}

	for k := 0; k < numVertices; k++ {
		for i := 0; i < numVertices; i++ {
			for j := 0; j < numVertices; j++ {
				if result[i][j] > result[i][k]+result[k][j] {
					result[i][j] = result[i][k] + result[k][j]
				}
			}
		}
	}

	return result

}
