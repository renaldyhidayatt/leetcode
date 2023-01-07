package main

func uniquePaths(m, n int) int {
	dmap := make([][]int, m)
	for i := range dmap {
		dmap[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dmap[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dmap[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			l, u := 0, 0
			if i-1 >= 0 {
				u = dmap[i-1][j]
			}
			if j-1 >= 0 {
				l = dmap[i][j-1]
			}
			dmap[i][j] = l + u
		}
	}
	return dmap[m-1][n-1]
}
