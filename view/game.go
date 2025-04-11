package view

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 700
	miniSize     = 400
)

type Game struct {
	miniImage   *ebiten.Image
	miniGameMap map[rune]*MiniGame
	startTime   time.Time
}

type MiniGame struct {
	alarmPos  float32
	miniImage *ebiten.Image
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.miniImage = g.miniGameMap['w'].miniImage
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.miniImage = g.miniGameMap['a'].miniImage
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.miniImage = g.miniGameMap['s'].miniImage
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.miniImage = g.miniGameMap['d'].miniImage
	} else if g.miniImage == nil {
		g.miniImage = ebiten.NewImage(miniSize, miniSize)
		g.miniImage.Fill(color.RGBA{0, 0, 0, 255})
	}
	// need to do behavoir for each of the images/minigames
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw backgroud
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff}) // change this to a custom image
	// draw game alerts
	for _, miniGame := range g.miniGameMap {
		vector.DrawFilledRect(screen, miniGame.alarmPos, 50, 75, 75, color.RGBA{0, 255, 0, 255}, false)
	}
	// draw mini image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64((ScreenWidth-miniSize)/2), float64(ScreenHeight-miniSize)-100)
	screen.DrawImage(g.miniImage, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	game := Game{
		miniGameMap: make(map[rune]*MiniGame, 4),
		startTime:   time.Now(),
	}
	game.miniGameMap['w'] = &MiniGame{
		alarmPos:  25,
		miniImage: ebiten.NewImage(miniSize, miniSize),
	}
	game.miniGameMap['a'] = &MiniGame{
		alarmPos:  125,
		miniImage: ebiten.NewImage(miniSize, miniSize),
	}
	game.miniGameMap['s'] = &MiniGame{
		alarmPos:  600,
		miniImage: ebiten.NewImage(miniSize, miniSize),
	}
	game.miniGameMap['d'] = &MiniGame{
		alarmPos:  700,
		miniImage: ebiten.NewImage(miniSize, miniSize),
	}

	return &game
}
