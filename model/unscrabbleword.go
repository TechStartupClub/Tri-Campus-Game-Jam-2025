package model

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type UnscrabbleWord struct {
	BaseMiniGame
}

func (game *UnscrabbleWord) UpdateGame() {
	game.MiniImage.Fill(color.Black)
	ebitenutil.DebugPrint(game.MiniImage, "Unscrabble Word")
}
