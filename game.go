package main

import (
	_ "image/png"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Fruit struct {
	x, y     int
	image    *ebiten.Image
	v0, hMax float64
	state    string
}

func (f *Fruit) MoveUp() {
	s := math.Sqrt(f.hMax+float64(f.y))/2 + f.v0
	f.y -= int(s)
}

func (f *Fruit) MoveDown() {
	s := math.Sqrt(f.hMax+float64(f.y)) / 4
	f.y += int(s)
}

type Game struct {
	fruits    []*Fruit
	fruitsImg []*ebiten.Image
	amount    int
}

func (g *Game) DefineParams() {
	g.fruitsImg = LoadImgs()
	g.amount = 10
	g.fruits = make([]*Fruit, g.amount)

	for i := 0; i < g.amount; i++ {
		randomIdx := rand.Intn(16)
		g.fruits[i] = &Fruit{
			x:     rand.Intn(WIDTH-100) + 10,
			y:     HEIGHT + 10,
			image: g.fruitsImg[randomIdx],
			v0:    float64(rand.Intn(4) + 1),
			hMax:  float64(rand.Intn(HEIGHT-200) + 100),
			state: "up",
		}
	}
}

func (g *Game) Update() {
	for _, fruit := range g.fruits {
		if fruit.y > int(fruit.hMax) && fruit.state == "up" {
			fruit.MoveUp()
		} else if fruit.y <= int(fruit.hMax) && fruit.state == "up" {
			fruit.state = "down"
		} else if fruit.y < (HEIGHT+10) && fruit.state == "down" {
			fruit.MoveDown()
		} else if fruit.y >= (HEIGHT+10) && fruit.state == "down" {
			fruit.state = "up"

		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, fruit := range g.fruits {
		op := ChangePos(fruit.x, fruit.y)
		screen.DrawImage(fruit.image, op)
	}
}
