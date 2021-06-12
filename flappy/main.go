package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	screenWidth   = 640
	screenHeight  = 640
	titleSize     = 32
	titleFontSize = fontSize * 1.5
	fontSize      = 24
	smallFontSize = fontSize / 2
)

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

func isAnyKeyJustPressed() bool {
	for _, k := range inpututil.PressedKeys() {
		if inpututil.IsKeyJustPressed(k) {
			return true
		}
	}
	return false

}
func jump() bool {
	if isAnyKeyJustPressed() {
		return true
	}
	return false
}

type Game struct {
	mode    Mode
	x16     int
	y16     int
	cameraX int
	cameraY int

	pipeTileYs []int
}

func (g *Game) init() {
	g.x16 = 0
	g.y16 = 100 * 16
	g.cameraX = -240
	g.cameraY = 0
	g.pipeTileYs = make([]int, 256)
	for i := range g.pipeTileYs {
		g.pipeTileYs[i] = rand.Intn(6) + 2
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return width, height
}

func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if jump() {
			g.mode = ModeGame
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	g.drawTiles(screen)
	if g.mode != ModeTitle {
	}
	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"FLappy"}
		texts = []string{"", "", "", "", "", "", "", "Press any key", "", "or touch screen"}
	}
	fmt.Println(screen)
	for i, l := range titleTexts {
		x := (screenWidth - len(l)*titleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*titleFontSize, color.White)
	}
	for i, l := range texts {
		x := (screenWidth - len(l)*fontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*fontSize, color.White)
	}

}

func (g *Game) drawTiles(screen *ebiten.Image) {

}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Flappy Demo")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
