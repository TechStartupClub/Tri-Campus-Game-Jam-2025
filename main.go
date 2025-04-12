package main

import (
	"log"

	"Tri-Campus-Game-Jam-2025/view"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := view.NewGame()

	ebiten.SetWindowSize(view.ScreenWidth, view.ScreenHeight)
	ebiten.SetWindowTitle("Avoid the Inevitable")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
