package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if m == 0 {
		return 0
	}
	dmap := make([][]int, m+1)
	for i := range dmap {
		dmap[i] = make([]int, n+1)
	}
	dmap[m-1][n] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				dmap[i][j] = 0
			} else {
				dmap[i][j] = dmap[i][j+1] + dmap[i+1][j]
			}
		}
	}
	return dmap[0][0]
}
