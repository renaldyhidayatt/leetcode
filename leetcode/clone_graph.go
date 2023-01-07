package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	labelMap := make(map[int]*Node)
	queue := []*Node{node}
	graphCopy := &Node{Val: node.Val, Neighbors: nil}

	labelMap[node.Val] = graphCopy

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, ne := range curr.Neighbors {
			if _, ok := labelMap[ne.Val]; ok {
				labelMap[curr.Val].Neighbors = append(labelMap[curr.Val].Neighbors, labelMap[ne.Val])

			} else {
				neighborCopy := &Node{Val: ne.Val, Neighbors: nil}
				labelMap[curr.Val].Neighbors = append(labelMap[curr.Val].Neighbors, neighborCopy)
				labelMap[ne.Val] = neighborCopy
				queue = append(queue, ne)
			}
		}
	}

	return graphCopy
}
