package stateoftictactoe

import "fmt"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

const PlayerO byte = 'O'
const PlayerX byte = 'X'

func StateOfTicTacToe(board []string) (State, error) {
	OCount := CountPlayer(board, PlayerO)
	XCount := CountPlayer(board, PlayerX)

	if !(OCount == XCount || XCount == OCount+1) {
		return "", fmt.Errorf("game board invalid")
	}

	OWin := WinPlayer(board, PlayerO)
	XWin := WinPlayer(board, PlayerX)

	if OWin && XWin {
		return "", fmt.Errorf("game board invalid")
	}

	if OWin || XWin {
		return Win, nil
	}

	if OCount+XCount == 9 {
		return Draw, nil
	}

	return Ongoing, nil
}

func CountPlayer(board []string, player byte) int {
	count := 0
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if player == board[r][c] {
				count += 1
			}
		}
	}
	return count
}

func WinPlayer(board []string, player byte) bool {
	for i := 0; i < 3; i++ {
		if player == board[i][0] && player == board[i][1] && player == board[i][2] {
			return true
		}
		if player == board[0][i] && player == board[1][i] && player == board[2][i] {
			return true
		}
	}
	if player == board[0][0] && player == board[1][1] && player == board[2][2] {
		return true
	}
	if player == board[2][0] && player == board[1][1] && player == board[0][2] {
		return true
	}
	return false
}
