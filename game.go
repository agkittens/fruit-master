package main

import (
	_ "image/png"
	"math"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type FlyingObj struct {
	x, y            int
	image           *ebiten.Image
	v0, hMax, theta float64
	state           string
	sX, sYU, sYD    float64
	slowFactor      float64
}

func (f *FlyingObj) DefineConsts() {
	side := 1.0
	if rand.Intn(2) == 0 {
		side = -1.0
	}
	f.sX = side * f.v0 * math.Cos(f.theta)
	f.sYU = math.Sqrt(f.hMax+float64(f.y))/2 + f.v0*math.Sin(f.theta)
	f.sYD = math.Sqrt(f.hMax+float64(f.y))/3 + f.v0*math.Sin(f.theta)
}
func (f *FlyingObj) MoveUp() {
	f.slowFactor = 1.0 - float64(f.y)/float64(HEIGHT-50)

	if f.slowFactor < 0.1 {
		f.slowFactor = 0.1
	}
	f.x -= int(f.sX)
	f.y -= (int(f.sYU)*int(f.slowFactor) + int(f.y/100))
}

func (f *FlyingObj) MoveDown() {
	f.slowFactor = 1.0 - float64(f.y)/float64(HEIGHT-50)
	if f.slowFactor < 0.1 {
		f.slowFactor = 0.1
	}
	f.x -= int(f.sX)
	f.y += (int(f.sYD)*int(f.slowFactor) + int(f.y/100))
}

func (f *FlyingObj) SmashObj() bool {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		posX, posY := ebiten.CursorPosition()

		if (posX >= f.x && posX <= f.x+f.image.Bounds().Dx()) && (posY >= f.y && posY <= f.y+f.image.Bounds().Dy()) {
			return true
		}
	}
	return false
}

func (f *FlyingObj) Movement() bool {
	if f.y > int(f.hMax) && f.state == "up" {
		f.MoveUp()
	} else if f.y <= int(f.hMax) && f.state == "up" {
		f.state = "down"
	} else if f.y < (HEIGHT+10) && f.state == "down" {
		f.MoveDown()
	} else if f.y >= (HEIGHT+10) && f.state == "down" {
		return true
	}
	return false
}

type Game struct {
	fruits    []*FlyingObj
	fruitsImg []*ebiten.Image
	bombs     *FlyingObj
	bombImg   []*ebiten.Image
	amount    int
	count     int
}

func (g *Game) DefineParams() {
	g.fruitsImg = LoadImgs(PATH_F)
	g.bombImg = LoadImgs(PATH_B)
	g.amount = 4
	g.fruits = make([]*FlyingObj, g.amount)
	g.count = 0

	for i := 0; i < g.amount; i++ {
		g.fruits[i] = g.CreateFruit()
	}

	g.bombs = &FlyingObj{
		x:     rand.Intn(WIDTH-100) + 10,
		y:     HEIGHT + 10,
		image: g.bombImg[0],
		v0:    float64(rand.Intn(4) + 1),
		hMax:  float64(rand.Intn(HEIGHT-200) + 100),
		theta: 20.0 * (math.Pi / 180),
		state: "up",
	}
	g.bombs.DefineConsts()
}

func (g *Game) Update() {
	for i := 0; i < len(g.fruits); i++ {
		fruit := g.fruits[i]
		changeStatus := fruit.Movement()

		if fruit.SmashObj() || changeStatus {
			g.fruits = append(g.fruits[:i], g.fruits[i+1:]...)
			g.fruits = append(g.fruits, g.CreateFruit())
			if fruit.SmashObj() {
				g.count += 1
			}
		}
	}
	g.bombs.Movement()
	if g.bombs.SmashObj() {
		g.count = 0
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, fruit := range g.fruits {
		op := ChangePos(fruit.x, fruit.y)
		screen.DrawImage(fruit.image, op)
	}
	op := ChangePos(g.bombs.x, g.bombs.y)
	screen.DrawImage(g.bombs.image, op)
	op = ChangePos(-10, -10)
	screen.DrawImage(g.fruitsImg[31], op)
	DisplayText(110, 30, 42, strconv.Itoa(g.count), screen)
}

func (g *Game) CreateFruit() *FlyingObj {
	randomIdx := rand.Intn(30)
	fruit := &FlyingObj{
		x:     rand.Intn(WIDTH-100) + 10,
		y:     HEIGHT + 10,
		image: g.fruitsImg[randomIdx],
		v0:    float64(rand.Intn(4) + 1),
		hMax:  float64(rand.Intn(HEIGHT-200) + 100),
		theta: 20.0 * (math.Pi / 180),
		state: "up",
	}
	fruit.DefineConsts()
	return fruit
}
