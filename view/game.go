package view

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 700
)

type Game struct {
	miniA *MiniGame
	miniB *MiniGame
	miniC *MiniGame
	miniD *MiniGame
}

type MiniGame struct {
	alarmPos float32
	gameType int
}

func (g *Game) Update() error {
	return nil
}

func DrawMiniGames(g *Game, screen *ebiten.Image) {
	vector.DrawFilledRect(screen, g.miniA.alarmPos, 50, 75, 75, color.RGBA{255, 0, 0, 255}, false)
	vector.DrawFilledRect(screen, g.miniB.alarmPos, 50, 75, 75, color.RGBA{255, 0, 0, 255}, false)
	vector.DrawFilledRect(screen, g.miniC.alarmPos, 50, 75, 75, color.RGBA{255, 0, 0, 255}, false)
	vector.DrawFilledRect(screen, g.miniD.alarmPos, 50, 75, 75, color.RGBA{255, 0, 0, 255}, false)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	DrawMiniGames(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	game := Game{
		miniA: &MiniGame{
			alarmPos: 25,
			gameType: 1,
		},
		miniB: &MiniGame{
			alarmPos: 125,
			gameType: 2,
		},
		miniC: &MiniGame{
			alarmPos: 600,
			gameType: 3,
		},
		miniD: &MiniGame{
			alarmPos: 700,
			gameType: 4,
		},
	}
	return &game
}
