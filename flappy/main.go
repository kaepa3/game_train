package main

import (
	_ "image/png"

	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math/rand"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	resources "github.com/hajimehoshi/ebiten/v2/examples/resources/images/flappy"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func floorDiv(x, y int) int {
	d := x / y
	if d*y == x || x >= 0 {
		return d
	}
	return d - 1
}

func floorMod(x, y int) int {
	return x - floorDiv(x, y)*y
}

const (
	screenWidth   = 640
	screenHeight  = 640
	tileSize      = 32
	titleFontSize = fontSize * 1.5
	fontSize      = 24
	smallFontSize = fontSize / 2
)

var (
	gopherImage     *ebiten.Image
	tilesImage      *ebiten.Image
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

func init() {
	fmt.Println("init-1")
	img, _, err := image.Decode(bytes.NewReader(resources.Gopher_png))
	if err != nil {
		log.Fatal(err)
	}
	gopherImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(resources.Tiles_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}
func init() {
	fmt.Println("init-2")
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    titleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

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
	const (
		nx           = screenWidth / tileSize
		ny           = screenHeight / tileSize
		pipeTileSrcX = 128
		pipeTileSrcY = 192
	)
	op := &ebiten.DrawImageOptions{}
	for i := -2; i < nx+1; i++ {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(i*tileSize-floorMod(g.cameraX, tileSize)),
			float64((ny-1)*tileSize-floorMod(g.cameraY, tileSize)))
		screen.DrawImage(tilesImage.SubImage(image.Rect(0, 0, tileSize, tileSize)).(*ebiten.Image), op)

	}

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
