package model

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func GetCurrNum() int {
	n := -1
	if ebiten.IsKeyPressed(ebiten.KeyDigit1) {
		n = 1
	} else if ebiten.IsKeyPressed(ebiten.KeyDigit2) {
		n = 2
	} else if ebiten.IsKeyPressed(ebiten.KeyDigit2) {
		n = 3
	} else if ebiten.IsKeyPressed(ebiten.KeyDigit2) {
		n = 4
	} else if ebiten.IsKeyPressed(ebiten.KeyDigit2) {
		n = 5
	} else if ebiten.IsKeyPressed(ebiten.KeyDigit2) {
		n = 6
	} else if ebiten.IsKeyPressed(ebiten.KeyDigit2) {
		n = 7
	} else if ebiten.IsKeyPressed(ebiten.KeyDigit2) {
		n = 8
	} else if ebiten.IsKeyPressed(ebiten.KeyDigit2) {
		n = 9
	}
	return n
}
