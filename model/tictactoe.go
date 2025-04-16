package model

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TicTacToeGame struct {
	BaseMiniGame
	board [3][3]int
}

func (game *TicTacToeGame) UpdateGame() {
	game.MiniImage.Fill(color.Black)

	msg := fmt.Sprintf("TicTacToe Game")
	op := &text.DrawOptions{}
	op.GeoM.Translate(20, 20)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(game.MiniImage, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   24,
	}, op)
}
