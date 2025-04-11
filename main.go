package main

import (
	"log"

	"GameJam/view"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := view.NewGame()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
