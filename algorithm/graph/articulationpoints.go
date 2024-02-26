package graph

import (
	"algoritmAndDs/graph/coloring"
	"algoritmAndDs/math/min"
)

type apHelper struct {
	is_ap              []bool
	visited            []bool
	child_cnt          []int
	discovery_time     []int
	earliest_discovery []int
}

func ArticulationPoint(graph *coloring.Graph) []bool {
	time := 0

	apHelperInstance := &apHelper{
		is_ap:              make([]bool, graph.Vertices),
		visited:            make([]bool, graph.Vertices),
		child_cnt:          make([]int, graph.Vertices),
		discovery_time:     make([]int, graph.Vertices),
		earliest_discovery: make([]int, graph.Vertices),
	}

	articulationPointHelper(
		apHelperInstance,
		0,
		-1,
		&time,
		graph,
	)

	if apHelperInstance.child_cnt[0] == 1 {
		apHelperInstance.is_ap[0] = false
	}

	return apHelperInstance.is_ap
}

func articulationPointHelper(apHelperInstance *apHelper, vertex int, parent int, time *int, graph *coloring.Graph) {
	apHelperInstance.visited[vertex] = true

	apHelperInstance.discovery_time[vertex] = *time
	apHelperInstance.earliest_discovery[vertex] = apHelperInstance.discovery_time[vertex]
	*time++

	for next_vertex := range graph.Edges[vertex] {
		if next_vertex == parent {
			continue
		}

		if apHelperInstance.visited[next_vertex] {
			apHelperInstance.earliest_discovery[vertex] = min.Int(
				apHelperInstance.earliest_discovery[vertex],
				apHelperInstance.discovery_time[next_vertex],
			)

			continue
		}

		apHelperInstance.child_cnt[vertex]++
		articulationPointHelper(
			apHelperInstance,
			next_vertex,
			vertex,
			time,
			graph,
		)

		apHelperInstance.earliest_discovery[vertex] = min.Int(
			apHelperInstance.earliest_discovery[vertex],
			apHelperInstance.earliest_discovery[next_vertex],
		)

		if apHelperInstance.earliest_discovery[next_vertex] >= apHelperInstance.discovery_time[vertex] {
			apHelperInstance.is_ap[vertex] = true
		}
	}
}
