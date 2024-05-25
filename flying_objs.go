package main

import (
	_ "image/png"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type FlyingObj struct {
	x, y            int
	image           *ebiten.Image
	v0, hMax, theta float64
	state           string
	sX, sY          float64
	slowFactor      float64
}

func (f *FlyingObj) DefineConsts() {
	side := 1.0
	if rand.Intn(2) == 0 {
		side = -1.0
	}
	f.sX = side * f.v0 * math.Cos(f.theta)
	f.sY = math.Sqrt(f.hMax+float64(f.y))/3 + f.v0*math.Sin(f.theta)
}
func (f *FlyingObj) MoveUp() {
	f.AdjustSlowFactor()

	f.x -= int(f.sX)
	f.y -= int((f.sY)*(f.slowFactor) + float64(f.y/100))

	if f.y <= int(f.hMax) {
		f.state = "down"
	}
}

func (f *FlyingObj) MoveDown() {
	f.AdjustSlowFactor()

	f.x -= int(f.sX)
	f.y += int((f.sY)*(f.slowFactor) + float64(f.y/100))
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

func (f *FlyingObj) AdjustSlowFactor() {
	f.slowFactor = float64(f.y)/float64(f.hMax) - 1.0
	if f.slowFactor < 0.01 {
		f.slowFactor = 0.01
	} else if f.slowFactor > 0.99 {
		f.slowFactor = 0.99
	}
}

func (f *FlyingObj) Move() {
	switch f.state {
	case "up":
		f.MoveUp()
	case "down":
		f.MoveDown()
	}
}

func (f *FlyingObj) CheckPos() bool {
	if f.y >= (HEIGHT+10) && f.state == "down" {
		return true
	}
	return false

}
