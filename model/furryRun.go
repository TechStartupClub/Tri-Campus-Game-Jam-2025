package model

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	numLanes     = 4
	laneHeight   = 20
	gridSize     = 32
	playerRadius = 10
)

type FurryRun struct {
	BaseMiniGame
	player     you
	cars       []car
	logs       []logThing
	laneColors []color.RGBA
}

type location struct {
	x float32
	y float32
}

type you struct {
	location
	width float32
}

type car struct {
	location
	width float32
	color color.RGBA
}

type logThing struct {
	location
	width float32
	color color.RGBA
}

func (game *FurryRun) Init() {
	game.player.location.x = ScreenWidth/2 + playerRadius
	game.player.location.y = ScreenHeight - 10

	game.laneColors = []color.RGBA{
		{100, 100, 100, 255}, // Road
		{100, 100, 100, 255}, // Road
		{50, 200, 50, 255},   // Grass
		{0, 0, 200, 255},     // Water
		{0, 0, 200, 255},     // Water
	}

	for i := range game.cars {
		lane := rand.Intn(2)
		game.cars[i] = car{
			location: location{
				x: float32(rand.Intn(ScreenWidth)),
				y: float32(lane+1)*gridSize + gridSize/2,
			},
			width: 64,
			color: color.RGBA{uint8(rand.Intn(200) + 55), uint8(rand.Intn(200) + 55), uint8(rand.Intn(200) + 55), 255},
		}
	}

	for i := range game.logs {
		lane := rand.Intn(2) + 3
		game.logs[i] = logThing{
			location: location{
				x: float32(rand.Intn(ScreenWidth)),
				y: float32(lane)*gridSize + gridSize/2,
			},
			width: 96,
		}
	}
}

// update mini image based on game state
func (game *FurryRun) UpdateScreen() {
	screen := game.MiniImage
	screen.Fill(color.RGBA{0, 150, 0, 255})

	// lane
	laneHeight := ScreenHeight / numLanes
	for i, laneColor := range game.laneColors {
		vector.DrawFilledRect(screen, 0, float32(i*laneHeight), ScreenWidth, float32(laneHeight), laneColor, false)
	}

	// safe zone
	vector.DrawFilledRect(screen, 0, 0, ScreenWidth, gridSize, color.RGBA{100, 255, 50, 255}, false)

	// car
	for _, car := range game.cars {
		vector.DrawFilledRect(screen, car.x-car.width/2, car.y-gridSize/2, car.width, gridSize, color.Black, false)
		vector.DrawFilledCircle(screen, car.x+car.width/2-10, car.y-gridSize/4, 5, color.RGBA{255, 255, 0, 255}, false)
		vector.DrawFilledCircle(screen, car.x+car.width/2-10, car.y+gridSize/4, 5, color.RGBA{255, 255, 0, 255}, false)
	}

	// log
	for _, log := range game.logs {
		vector.DrawFilledRect(screen, log.x-log.width/2, log.y-gridSize/2, log.width, gridSize, color.RGBA{139, 69, 19, 255}, false)
		for i := 0; i < int(log.width)-10; i += 10 {
			vector.DrawFilledRect(screen, log.x-log.width/2+float32(i), log.y-gridSize/2, log.x-log.width/2+float32(i), log.y+gridSize/2,
				color.RGBA{101, 67, 33, 255}, false)
		}
	}

	// player
	vector.DrawFilledCircle(screen, float32(game.player.x), float32(game.player.y), playerRadius, color.RGBA{0, 255, 0, 255}, false)

	vector.DrawFilledCircle(screen, game.player.x-gridSize/6, game.player.y-gridSize/6, gridSize/8, color.RGBA{255, 255, 255, 255}, false)
	vector.DrawFilledCircle(screen, game.player.x+gridSize/6, game.player.y-gridSize/6, gridSize/8, color.RGBA{255, 255, 255, 255}, false)
	vector.DrawFilledCircle(screen, game.player.x-gridSize/6, game.player.y-gridSize/6, gridSize/16, color.RGBA{0, 0, 0, 255}, false)
	vector.DrawFilledCircle(screen, game.player.x+gridSize/6, game.player.y-gridSize/6, gridSize/16, color.RGBA{0, 0, 0, 255}, false)

	vector.StrokeLine(screen, game.player.x-gridSize/4, game.player.y-gridSize/4, game.player.x-gridSize/2,
		game.player.y-gridSize/2, float32(2), color.RGBA{0, 200, 0, 255}, false)
	vector.StrokeLine(screen, game.player.x+gridSize/4, game.player.y-gridSize/4, game.player.x+gridSize/2,
		game.player.y-gridSize/2, float32(2), color.RGBA{0, 200, 0, 255}, false)
	vector.StrokeLine(screen, game.player.x-gridSize/4, game.player.y+gridSize/4, game.player.x-gridSize/2,
		game.player.y+gridSize/2, float32(2), color.RGBA{0, 200, 0, 255}, false)
	vector.StrokeLine(screen, game.player.x+gridSize/4, game.player.y+gridSize/4, game.player.x+gridSize/2,
		game.player.y+gridSize/2, float32(2), color.RGBA{0, 200, 0, 255}, false)

	ebitenutil.DebugPrint(screen, "Frogger")
}

func (game *FurryRun) UpdateGame() {
	// frog move up down left right and need to avoid water and car
}

func (game *FurryRun) checkWin() bool {
	// frog reach top, good, other wise not good
	return false
}

func (game *FurryRun) Reset() {
	Init()
}
