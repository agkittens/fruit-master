package main

import (
	_ "image/png"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Fruit struct {
	x, y     int
	image    *ebiten.Image
	v0, hMax float64
	state    string
	// onClick  func()
}

func (f *Fruit) MoveUp() {
	s := math.Sqrt(f.hMax+float64(f.y))/2 + f.v0
	f.y -= int(s)
}

func (f *Fruit) MoveDown() {
	s := math.Sqrt(f.hMax+float64(f.y)) / 4
	f.y += int(s)
}

func (f *Fruit) SmashFruit() bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		posX, posY := ebiten.CursorPosition()

		if (posX >= f.x && posX <= f.x+f.image.Bounds().Dx()) && (posY >= f.y && posY <= f.y+f.image.Bounds().Dy()) {
			// f.onClick()
			return true
		}
	}
	return false
}

type Game struct {
	fruits    []*Fruit
	fruitsImg []*ebiten.Image
	amount    int
	count     int
}

func (g *Game) DefineParams() {
	g.fruitsImg = LoadImgs()
	g.amount = 10
	g.fruits = make([]*Fruit, g.amount)
	g.count = 0

	for i := 0; i < g.amount; i++ {
		g.fruits[i] = g.CreateFruit()
	}
}

func (g *Game) Update() {
	for i := 0; i < len(g.fruits); i++ {
		fruit := g.fruits[i]
		if fruit.y > int(fruit.hMax) && fruit.state == "up" {
			fruit.MoveUp()
		} else if fruit.y <= int(fruit.hMax) && fruit.state == "up" {
			fruit.state = "down"
		} else if fruit.y < (HEIGHT+10) && fruit.state == "down" {
			fruit.MoveDown()
		} else if fruit.y >= (HEIGHT+10) && fruit.state == "down" {
			g.ChangeParams(fruit)
		}

		if fruit.SmashFruit() {
			g.count += 1
			log.Println("smash count", g.count)
			g.fruits = append(g.fruits[:i], g.fruits[i+1:]...)
			g.fruits = append(g.fruits, g.CreateFruit())
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, fruit := range g.fruits {
		op := ChangePos(fruit.x, fruit.y)
		screen.DrawImage(fruit.image, op)
	}
}

func (g *Game) ChangeParams(fruit *Fruit) {
	randomIdx := rand.Intn(30)
	fruit.x = rand.Intn(WIDTH-100) + 10
	fruit.y = HEIGHT + 10
	fruit.image = g.fruitsImg[randomIdx]
	fruit.v0 = float64(rand.Intn(4) + 1)
	fruit.hMax = float64(rand.Intn(HEIGHT-200) + 100)
	fruit.state = "up"

}

func (g *Game) CreateFruit() *Fruit {
	randomIdx := rand.Intn(30)
	fruit := &Fruit{
		x:     rand.Intn(WIDTH-100) + 10,
		y:     HEIGHT + 10,
		image: g.fruitsImg[randomIdx],
		v0:    float64(rand.Intn(4) + 1),
		hMax:  float64(rand.Intn(HEIGHT-200) + 100),
		state: "up",
	}
	return fruit
}
