package model

import (
	"bytes"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var mplusFaceSource *text.GoTextFaceSource

type MiniGame interface {
	UpdateGame()
	UpdateScreen()
	SetActive()
	Reset()
	IsActive() bool
	GetAlarmPos() float32
	GetMiniImage() *ebiten.Image
}

type BaseMiniGame struct {
	AlarmPos  float32
	MiniImage *ebiten.Image
	Active    bool
}

func Init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

// should update game state
func (game *BaseMiniGame) UpdateGame() {
	// this should just be empty
	// should be overridden
}

// update MiniImage based on game state
func (game *BaseMiniGame) UpdateScreen() {
	// this should just be empty
	// should be overridden
}

// reset to 'og' game state
func (game *BaseMiniGame) Reset() {
	// this should be empty
	// should be overriden
}

func (game *BaseMiniGame) SetActive() {
	game.Reset()
	game.Active = true
}

func (game *BaseMiniGame) IsActive() bool {
	return game.Active
}

func (game *BaseMiniGame) GetAlarmPos() float32 {
	return game.AlarmPos
}

func (game *BaseMiniGame) GetMiniImage() *ebiten.Image {
	return game.MiniImage
}
