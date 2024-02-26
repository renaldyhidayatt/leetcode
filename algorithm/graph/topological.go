package graph

func Topological(N int, constraint [][]int) []int {
	dependencies := make([]int, N)
	nodes := make([]int, N)

	for i := range nodes {
		nodes[i] = i
	}

	edges := make([][]bool, N)

	for i := range edges {
		edges[i] = make([]bool, N)
	}

	for _, c := range constraint {
		a := c[0]
		b := c[1]
		dependencies[b]++
		edges[a][b] = true
	}

	var answer []int

	for s := 0; s < N; s++ {
		if dependencies[s] == 0 {
			route, _ := DepthFirstSearchHelper(s, N, nodes, edges, true)
			answer = append(answer, route...)
		}
	}

	return answer
}
