package model

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type CopyNumber struct {
	BaseMiniGame
}

func (game *CopyNumber) UpdateGame() {
	game.MiniImage.Fill(color.Black)
	ebitenutil.DebugPrint(game.MiniImage, "CopyNumber")
}
