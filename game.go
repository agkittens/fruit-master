package main

import (
	"image/color"
	_ "image/png"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameData struct {
	Count  int    `json:"count"`
	Player string `json:"player"`
}

type Game struct {
	fruits        []*FlyingObj
	fruitsImg     []*ebiten.Image
	icon          *ebiten.Image
	bombs         *FlyingObj
	bombImg       []*ebiten.Image
	amount        int
	count         int
	hearts        int
	lastSpawnTime time.Time
	intensityTime time.Time
	gameOverTime  time.Time
	particles     []*Particles
	gameOver      bool
}

func (g *Game) DefineParams() {
	g.fruitsImg = LoadImgs(PATH_F)
	g.bombImg = LoadImgs(PATH_B)
	g.icon = g.fruitsImg[39]

	g.fruits = make([]*FlyingObj, g.amount)
	for i := 0; i < g.amount; i++ {
		g.fruits[i] = g.CreateFlyingObj("fruit", g.fruitsImg)
	}

	g.bombs = g.CreateFlyingObj("bomb", g.bombImg)
	g.lastSpawnTime = time.Now()
	g.intensityTime = time.Now()
	g.gameOver = false
}

func (g *Game) Update() {
	if g.hearts <= 0 {
		if !g.gameOver {
			g.gameOver = true
			g.gameOverTime = time.Now()
			gameData := &GameData{
				Count:  g.count,
				Player: "player1",
			}
			SaveGameData("data.json", gameData)
		}
		if time.Since(g.gameOverTime) >= 4*time.Second {
			currentState = StateScore
		}
		return
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.CreateParticle()
	}

	for i := 0; i < len(g.particles); i++ {
		g.particles[i].Fade()
	}

	g.ManageIntensity()
	for i := 0; i < len(g.fruits); i++ {
		fruit := g.fruits[i]
		fruit.Move()
		isFallenF := fruit.CheckPos()
		isSmashedF := fruit.SmashObj()
		if isSmashedF || isFallenF {
			g.fruits = append(g.fruits[:i], g.fruits[i+1:]...)
			g.fruits = append(g.fruits, g.CreateFlyingObj("fruit", g.fruitsImg))
			if isSmashedF {
				g.count++
			}

		}
	}
	g.bombs.Move()
	isFallenB := g.bombs.CheckPos()
	isSmashedB := g.bombs.SmashObj()
	if isFallenB && time.Since(g.lastSpawnTime) >= 20*time.Second {
		g.bombs = g.CreateFlyingObj("bomb", g.bombImg)
		g.lastSpawnTime = time.Now()
	}
	if isSmashedB {
		g.bombs = g.CreateFlyingObj("bomb", g.bombImg)
		g.hearts--
		if g.hearts > 0 {
			g.count = 0
		}
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.hearts <= 0 {
		DisplayText(WIDTH/2-100, HEIGHT/3, 52, "You lost", screen, color.Black)
	}

	for i := 1; i < len(g.particles); i++ {
		if g.particles[i].active {
			prev := g.particles[i-1]
			curr := g.particles[i]
			vector.StrokeLine(screen, prev.x, prev.y, curr.x, curr.y, 7, color.RGBA{255, 255, 255, 1}, true)
		}
	}
	for _, fruit := range g.fruits {
		op := ChangePos(fruit.x, fruit.y)
		screen.DrawImage(fruit.image, op)
	}

	op := ChangePos(g.bombs.x, g.bombs.y)
	screen.DrawImage(g.bombs.image, op)

	op = ChangePos(10, 10)
	screen.DrawImage(g.icon, op)
	DisplayText(100, 25, 42, strconv.Itoa(g.count), screen, color.White)
	DisplayText(WIDTH-150, 25, 36, ("HP:" + strconv.Itoa(g.hearts)), screen, color.White)
}

func (g *Game) CreateFlyingObj(obj string, arr []*ebiten.Image) *FlyingObj {
	randomIdx := 0
	vel := 6

	if obj == "fruit" {
		randomIdx = rand.Intn(30)
		vel = 4
	}

	object := &FlyingObj{
		x:     rand.Intn(WIDTH-100) + 100,
		y:     HEIGHT + 10,
		image: arr[randomIdx],
		v0:    float64(rand.Intn(vel) + 1),
		hMax:  float64(rand.Intn(HEIGHT-200) + 100),
		theta: 20.0 * (math.Pi / 180),
		state: "up",
	}
	object.DefineConsts()
	return object
}

func (g *Game) CreateParticle() {
	x, y := ebiten.CursorPosition()
	g.particles = append(g.particles,
		&Particles{
			x:         float32(x),
			y:         float32(y),
			alpha:     1.0,
			fadeSpeed: 0.1,
			active:    true})
}

func (g *Game) ManageIntensity() {
	if time.Since(g.intensityTime) >= 10*time.Second {
		if g.amount < 5 {
			g.amount++

			g.fruits = append(g.fruits, g.CreateFlyingObj("fruit", g.fruitsImg))
		}
		g.intensityTime = time.Now()
	}
}
