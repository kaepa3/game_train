package train

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	resources "github.com/kaepa3/game_train/train/resources"

	_ "image/png"
)

var (
	gopherImage    *ebiten.Image
	posX           float64
	posY           float64
	characterScale = 0.3
)

const (
	screenWidth  = 600
	screenHeight = 600
)

type Game struct{}

func Run() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("testing")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) init() {
	img, _, err := image.Decode(bytes.NewReader(resources.Gopher_png))
	if err != nil {
		log.Println(resources.Gopher_png)
		log.Fatal(err)
	}
	gopherImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	g.move()
	return nil
}

func (g *Game) move() {
	for _, k := range inpututil.PressedKeys() {
		if k == ebiten.KeyArrowUp {
			moveIfCan(0, -1)
		} else if k == ebiten.KeyArrowDown {
			moveIfCan(0, 1)
		} else if k == ebiten.KeyArrowRight {
			moveIfCan(1, 0)
		} else if k == ebiten.KeyArrowLeft {
			moveIfCan(-1, 0)
		}
	}
}
func moveIfCan(x, y float64) {
	w, h := gopherImage.Size()
	toX := posX + x
	toY := posY + y
	if toX > 0 && toX < screenWidth-(float64(w)*characterScale) {
		posX = toX
	}
	if toY > 0 && toY < screenHeight-(float64(h)*characterScale) {
		posY = toY
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawGopher(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %02.f", ebiten.CurrentFPS()))
}

func (g *Game) drawGopher(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(characterScale, characterScale)
	op.GeoM.Translate(posX, posY)
	screen.DrawImage(gopherImage, op)
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}
