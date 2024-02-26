package minesweeper

func minesAt(board []string, row, col int) int {
	mines := 0

	nrows := len(board)
	ncols := len(board[0])

	border := row + 1

	if border >= nrows {
		border = -1
	}

	rows := [8]int{
		row - 1,
		row - 1,
		row - 1,
		row,
		row,
		border,
		border,
		border,
	}

	border = col + 1

	if border >= ncols {
		border = -1
	}

	cols := [8]int{
		col - 1,
		col,
		border,
		col - 1,
		border,
		col - 1,
		col,
		border,
	}

	for i := 0; i < 8; i++ {
		r, c := rows[i], cols[i]

		if r >= 0 && c >= 0 && board[r][c] == '*' {
			mines++
		}
	}

	return mines
}

func Annotate(board []string) []string {
	nrows := len(board)

	if nrows == 0 {
		return []string{}
	}

	ncols := len(board[0])

	annotation := make([]string, 0, nrows)

	for r := 0; r < nrows; r++ {
		bb := []byte(board[r])

		for c := 0; c < ncols; c++ {
			if bb[c] == ' ' {
				mines := minesAt(board, r, c)

				if mines > 0 {
					bb[c] = '0' + byte(mines)
				}
			}
		}

		annotation = append(annotation, string(bb))
	}

	return annotation
}
