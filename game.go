package main

import (
	"fmt"
	_ "image/png"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	x, y        int
	fruitsImg   *ebiten.Image
	v0, g, hMax float64
	state       string
}

func (g *Game) DefineParams() {
	g.fruitsImg, _, _ = ebitenutil.NewImageFromFile(FB)
	g.x = rand.Intn(WIDTH-10) + 10
	g.y = HEIGHT + 10
	g.v0 = float64(rand.Intn(4) + 1)
	g.g = 9.81
	g.hMax = float64(rand.Intn(HEIGHT-200) + 100)
	fmt.Print(g.hMax, g.v0, g.y)
	g.state = "up"
}

func (g *Game) Update() {
	if g.y > int(g.hMax) && g.state == "up" {
		g.MoveUp()
	} else if g.y <= int(g.hMax) && g.state == "up" {
		g.state = "down"
	} else {
		g.MoveDown()
	}
}
func (g *Game) Draw(screen *ebiten.Image) {
	op := ChangePos(g.x, g.y)
	screen.DrawImage(g.fruitsImg, op)
}

func (g *Game) MoveUp() {
	s := math.Sqrt((g.hMax + float64(g.y))) / 2
	g.y -= int(s)
}

func (g *Game) MoveDown() {
	s := math.Sqrt((g.hMax + float64(g.y))) / 4
	g.y += int(s)
}
