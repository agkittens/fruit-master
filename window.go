package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	background *ebiten.Image
	title      *ebiten.Image
}

func (g *Game) Init() {
	g.background, _, _ = ebitenutil.NewImageFromFile(BG)
	g.title, _, _ = ebitenutil.NewImageFromFile(TITLE)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opBG := g.AdjustSize(g.background, 2, 2)
	opTT := g.AdjustSize(g.title, 2, 3)
	screen.DrawImage(g.background, opBG)
	screen.DrawImage(g.title, opTT)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}

func (g *Game) AdjustSize(img *ebiten.Image, divX int, divY int) *ebiten.DrawImageOptions {
	size := img.Bounds().Size()
	posX := (WIDTH - size.X) / divX
	posY := (HEIGHT - size.Y) / divY
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(posX), float64(posY))
	return op
}
