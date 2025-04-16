package view

import (
	"image"
	"image/color"
	_ "image/png"
	"os"

	"Tri-Campus-Game-Jam-2025/model"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	AlarmSize    = 75
	ScreenWidth  = 800
	ScreenHeight = 700
	MiniSize     = 400
	MiniPosX     = 200
	MiniPosY     = 200
)

type Game struct {
	currMini    rune
	miniGameMap map[rune]model.MiniGame
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

	if g.currMini != 0 {
		g.miniGameMap[g.currMini].UpdateGame()
		g.miniGameMap[g.currMini].UpdateScreen()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// draw backgroud
	op1 := &ebiten.DrawImageOptions{}
	bounds := g.backImage.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	op1.GeoM.Scale(float64(ScreenWidth)/float64(w), float64(ScreenHeight)/float64(h))
	screen.DrawImage(g.backImage, op1)
	// draw game alerts
	for _, miniGame := range g.miniGameMap {
		green, red := 0, 0
		if miniGame.IsActive() {
			red = 255
		} else {
			green = 255
		}
		vector.DrawFilledRect(screen, miniGame.GetAlarmPos(), 50, AlarmSize, AlarmSize, color.RGBA{uint8(red), uint8(green), 0, 255}, false)
	}
	// draw curr minigame
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(MiniPosX, MiniPosY)
	if g.currMini != 0 && g.miniGameMap[g.currMini].IsActive() {
		screen.DrawImage(g.miniGameMap[g.currMini].GetMiniImage(), op2)
	} else {
		temp_screen := ebiten.NewImage(MiniSize, MiniSize)
		temp_screen.Fill(color.Black)
		screen.DrawImage(temp_screen, op2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() *Game {
	game := Game{
		miniGameMap: map[rune]model.MiniGame{
			'w': &model.TicTacToeGame{
				BaseMiniGame: model.BaseMiniGame{
					AlarmPos:  25,
					MiniImage: ebiten.NewImage(MiniSize, MiniSize),
					Active:    false,
				},
			},
			'a': &model.CopyNumber{
				BaseMiniGame: model.BaseMiniGame{
					AlarmPos:  125,
					MiniImage: ebiten.NewImage(MiniSize, MiniSize),
					Active:    false,
				},
			},
			's': &model.SortList{
				BaseMiniGame: model.BaseMiniGame{
					AlarmPos:  600,
					MiniImage: ebiten.NewImage(MiniSize, MiniSize),
					Active:    false,
				},
			},
			'd': &model.UnscrabbleWord{
				BaseMiniGame: model.BaseMiniGame{
					AlarmPos:  700,
					MiniImage: ebiten.NewImage(MiniSize, MiniSize),
					Active:    false,
				},
			},
		},
	}

	file, _ := os.Open("sprites/background.png")
	defer file.Close()
	img, _, _ := image.Decode(file)
	game.backImage = ebiten.NewImageFromImage(img)

	model.Init()

	return &game
}
