package graph

import "sort"

type Vertex int

type Edge struct {
	Start  Vertex
	End    Vertex
	Weight int
}

func KruskalMST(n int, edges []Edge) ([]Edge, int) {
	var mst []Edge
	var cost int

	u := NewUnionFind(n)

	for i := 0; i < n; i++ {
		u.parent[i] = i
		u.size[i] = 1
	}

	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	for _, edge := range edges {
		if u.Find(int(edge.Start)) != u.Find(int(edge.End)) {
			mst = append(mst, edge)

			cost += edge.Weight

			u = u.Union(int(edge.Start), int(edge.End))
		}
	}

	return mst, cost
}
