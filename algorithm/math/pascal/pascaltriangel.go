package pascal

func GenerateTriangle(n int) [][]int {
	var triangle = make([][]int, n)

	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)

		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1]
		}
	}

	return triangle
}
