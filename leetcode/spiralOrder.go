package main

func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	return getSpiralOrder(matrix, 0, m-1, 0, n-1)
}

func getSpiralOrder(matrix [][]int, rStart, rEnd, cStart, cEnd int) []int {
	if rStart > rEnd || cStart > cEnd {
		return []int{}
	} else if rStart == rEnd {
		return matrix[rStart][cStart : cEnd+1]
	} else if cStart == cEnd {
		curr := make([]int, rEnd-rStart+1)
		for j := range curr {
			curr[j] = matrix[j+rStart][cEnd]
		}
		return curr
	}
	curr := matrix[rStart][cStart : cEnd+1]
	for j := rStart + 1; j < rEnd; j++ {
		curr = append(curr, matrix[j][cEnd])
	}
	curr = append(curr, reversee(matrix[rEnd][cStart:cEnd+1])...)

	for j := rEnd - 1; j > rStart; j-- {
		curr = append(curr, matrix[j][cStart])
	}
	res := append(curr, getSpiralOrder(matrix, rStart+1, rEnd-1, cStart+1, cEnd-1)...)
	return res
}

func reversee(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
