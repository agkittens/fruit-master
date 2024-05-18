package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var currentState int

type Window struct {
	background *ebiten.Image
	title      *ebiten.Image
	buttons    []*Button
}

func (w *Window) Init() {
	w.background, _, _ = ebitenutil.NewImageFromFile(BG)
	w.title, _, _ = ebitenutil.NewImageFromFile(TITLE)

	buttonStart := &Button{
		x:      (WIDTH - 200) / 2,
		y:      (HEIGHT + 150) / 2,
		width:  200,
		height: 50,
		text:   "start",
		onClick: func() {
			currentState = StateGame
			w.background, _, _ = ebitenutil.NewImageFromFile(GAME)
		},
	}

	buttonExit := &Button{
		x:      (WIDTH - 200) / 2,
		y:      (HEIGHT + 300) / 2,
		width:  200,
		height: 50,
		text:   "exit",
		onClick: func() {
			currentState = StateExit
		},
	}

	buttonX := &Button{
		x:      (WIDTH - 50),
		y:      20,
		width:  30,
		height: 30,
		text:   "X",
		onClick: func() {
			currentState = StateMenu
			w.background, _, _ = ebitenutil.NewImageFromFile(BG)
		},
	}

	w.buttons = []*Button{buttonStart, buttonExit, buttonX}

}

func (w *Window) Update() error {
	switch currentState {
	case StateMenu:
		w.buttons[0].Update()
		w.buttons[1].Update()
	case StateGame:
		w.buttons[2].Update()
	case StateExit:
		return ebiten.Termination
	}
	return nil
}

func (w *Window) Draw(screen *ebiten.Image) {
	switch currentState {
	case StateMenu:
		opBG := AdjustSize(w.background, 2, 2)
		opTT := AdjustSize(w.title, 2, 3)
		screen.DrawImage(w.background, opBG)
		screen.DrawImage(w.title, opTT)
		w.buttons[0].Draw(screen)
		w.buttons[1].Draw(screen)

	case StateGame:
		opBG := AdjustSize(w.background, 2, 2)
		screen.DrawImage(w.background, opBG)
		w.buttons[2].Draw(screen)
	}

}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}
