package view

import (
	"image"
	"os"

	_ "image/png"

	"Tri-Campus-Game-Jam-2025/model"

	"github.com/hajimehoshi/ebiten/v2"
)

type MiniGame = model.MiniGame

const (
	ScreenWidth  = 800
	ScreenHeight = 700
	MiniSize     = 400
)

type Game struct {
	currMini    rune
	miniGameMap map[rune]MiniGame
	backImage   *ebiten.Image
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.currMini = 'w'
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.currMini = 'a'
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.currMini = 's'
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.currMini = 'd'
	}
	g.miniGameMap[g.currMini].UpdateGame()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw backgroud
	op := &ebiten.DrawImageOptions{}
	bounds := g.backImage.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	op.GeoM.Scale(float64(ScreenWidth)/float64(w), float64(ScreenHeight)/float64(h))
	screen.DrawImage(g.backImage, op)
	// draw game alerts
	for _, miniGame := range g.miniGameMap {
		miniGame.DisplayAlert(screen)
	}
	// draw mini game screen
	g.miniGameMap[g.currMini].DisplayScreen(screen, float64(ScreenWidth-MiniSize)/2, float64(ScreenHeight-MiniSize)-100)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	game := Game{
		miniGameMap: map[rune]MiniGame{
			'w': &model.TicTacToeGame{
				BaseMiniGame: model.BaseMiniGame{
					AlarmPos:  25,
					MiniImage: ebiten.NewImage(MiniSize, MiniSize),
				},
			},
			'a': &model.CopyNumber{
				BaseMiniGame: model.BaseMiniGame{
					AlarmPos:  125,
					MiniImage: ebiten.NewImage(MiniSize, MiniSize),
				},
			},
			's': &model.SortList{
				BaseMiniGame: model.BaseMiniGame{
					AlarmPos:  600,
					MiniImage: ebiten.NewImage(MiniSize, MiniSize),
				},
			},
			'd': &model.UnscrabbleWord{
				BaseMiniGame: model.BaseMiniGame{
					AlarmPos:  700,
					MiniImage: ebiten.NewImage(MiniSize, MiniSize),
				},
			},
		},
	}
	game.currMini = 'w'

	file, _ := os.Open("sprites/background.png")
	defer file.Close()
	img, _, _ := image.Decode(file)
	game.backImage = ebiten.NewImageFromImage(img)

	model.Init()

	return &game
}
