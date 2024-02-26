package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	var c int

	for _, v := range cb[file] {
		if v {
			c++
		}
	}

	return c
}

func CountInRank(cb Chessboard, rank int) int {
	var c int

	if rank >= 1 && rank <= 8 {
		for _, v := range cb {
			if v[rank-1] {
				c++
			}
		}
	}
	return c
}

func CountAll(cb Chessboard) int {
	var c int

	for _, v := range cb {
		for range v {
			c++
		}
	}

	return c
}
