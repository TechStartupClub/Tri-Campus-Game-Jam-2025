package model

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SortList struct {
	BaseMiniGame
}

func (game *SortList) UpdateGame() {
	game.MiniImage.Fill(color.Black)
	ebitenutil.DebugPrint(game.MiniImage, "SortList")
}
