package graph

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(s int) UnionFind {
	parent := make([]int, s)
	size := make([]int, s)

	for k := 0; k < s; k++ {
		parent[k] = k
		size[k] = 1
	}

	return UnionFind{parent: parent, size: size}
}

func (u UnionFind) Find(q int) int {
	for q != u.parent[q] {
		q = u.parent[q]
	}
	return q
}

func (u UnionFind) Union(a, b int) UnionFind {
	rootP := u.Find(a)
	rootQ := u.Find(b)

	if rootP == rootQ {
		return u
	}

	if u.size[rootP] < u.size[rootQ] {
		u.parent[rootP] = rootQ
		u.size[rootQ] += u.size[rootP]
	} else {
		u.parent[rootQ] = rootP
		u.size[rootP] += u.size[rootQ]
	}

	return u
}
