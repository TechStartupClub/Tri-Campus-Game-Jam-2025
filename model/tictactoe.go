package model

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	TTTX         = 20.0
	TTTY         = 60.0
	TTTLineSpace = 28.0
)

type TicTacToeGame struct {
	BaseMiniGame
	Board [3][3]int
}

// update mini image based on game state
func (game *TicTacToeGame) UpdateScreen() {
	game.MiniImage.Fill(color.Black)
	for i, row := range game.Board {
		rowStr := ""
		for _, cell := range row {
			switch cell {
			case 0:
				rowStr += "* "
			case 1:
				rowStr += "O "
			case 2:
				rowStr += "X "
			}
		}
		op := &text.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		op.GeoM.Translate(TTTX, TTTY+float64(i)*TTTLineSpace)
		text.Draw(game.MiniImage, rowStr, &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   20,
		}, op)
	}
}

// update game state
// check if should switch Active to false (if solved)
func (game *TicTacToeGame) UpdateGame() {
	input := GetCurrNum()
	row := (input - 1) / 3
	col := (input - 1) % 3

	if input != -1 && game.Board[row][col] == 0 {
		game.Board[row][col] = 2
		game.oppMove()
	}

	winner := game.checkWinner()

	if winner {
		game.Active = false
	}

}

func (game *TicTacToeGame) oppMove() {
	empty := [][2]int{}
	for i, row := range game.Board {
		for j, val := range row {
			if val == 0 {
				empty = append(empty, [2]int{i, j})
			}
		}
	}
	if len(empty) != 0 {
		move := empty[rand.Intn(len(empty))]
		game.Board[move[0]][move[1]] = 1
	}
}

func (game *TicTacToeGame) checkWinner() bool {
	board := game.Board
	for i := range board {
		if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			if board[i][0] == 1 {
				game.Reset()
				return false
			} else if board[i][0] == 2 {
				return true
			}
		}
	}
	for j := range board {
		if board[0][j] != 0 && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
			if board[0][j] == 1 {
				game.Reset()
				return false
			} else if board[0][j] == 2 {
				return true
			}
		}
	}
	if board[0][0] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[0][0] == 1 {
			game.Reset()
			return false
		} else if board[0][0] == 2 {
			return true
		}
	}
	if board[0][2] != 0 && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		if board[0][2] == 1 {
			game.Reset()
			return false
		} else if board[0][2] == 2 {
			return true
		}
	}
	full := true
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 0 {
				full = false
				break
			}
		}
	}
	if full {
		game.Reset()
	}
	return false
}

func (game *TicTacToeGame) Reset() {
	for i := range game.Board {
		for j := range game.Board[i] {
			game.Board[i][j] = 0
		}
	}
}
