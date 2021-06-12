package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "hello world")
}

func (g *Game) Layout(width, height int) (int, int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 240)
	ebiten.SetWindowTitle("hello world")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
