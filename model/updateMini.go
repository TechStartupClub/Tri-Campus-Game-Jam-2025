package model

import (
	"bytes"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var mplusFaceSource *text.GoTextFaceSource

type MiniGame interface {
	UpdateGame()
	DisplayAlert(*ebiten.Image)
	DisplayScreen(*ebiten.Image, float64, float64)
	GetMiniImage() *ebiten.Image
}

type BaseMiniGame struct {
	AlarmPos  float32
	MiniImage *ebiten.Image
}

func Init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

func (game *BaseMiniGame) GetMiniImage() *ebiten.Image {
	return game.MiniImage
}

func (game *BaseMiniGame) DisplayAlert(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, game.AlarmPos, 50, 75, 75, color.RGBA{0, 255, 0, 255}, false)
}

func (game *BaseMiniGame) DisplayScreen(screen *ebiten.Image, transX float64, transY float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(transX, transY)
	screen.DrawImage(game.MiniImage, op)
}
