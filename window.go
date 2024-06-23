package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var currentState int

type Window struct {
	background  *ebiten.Image
	scoreboard  *ebiten.Image
	title       *ebiten.Image
	gameplay    *Game
	playTicker  *time.Ticker
	pauseTicker *time.Ticker
	player      *audio.Player
	buttons     []*Button
}

func (w *Window) Init() {
	var err error
	w.background, _, err = ebitenutil.NewImageFromFile(BG)
	if err != nil {
		panic(err)
	}
	w.title, _, err = ebitenutil.NewImageFromFile(TITLE)
	if err != nil {
		panic(err)
	}
	w.gameplay = &Game{amount: 1, count: 0, hearts: 3, isMusic: false}

	buttonStart := &Button{
		x:      (WIDTH-200)/2 + 10,
		y:      (HEIGHT + 150) / 2,
		width:  172,
		height: 50,
		scaleX: 1,
		scaleY: 1,
		text:   "start",
		onClick: func() {
			currentState = StateGame
			w.background, _, err = ebitenutil.NewImageFromFile(GAME)
			if err != nil {
				panic(err)
			}
			w.scoreboard, _, err = ebitenutil.NewImageFromFile(SCOREBOARD)
			if err != nil {
				panic(err)
			}
			w.gameplay.ResetGame()
			w.PlayAudio()
		},
	}

	buttonExit := &Button{
		x:      (WIDTH-200)/2 + 10,
		y:      (HEIGHT + 300) / 2,
		width:  172,
		height: 50,
		scaleX: 1,
		scaleY: 1,
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
		scaleX: 0.2,
		scaleY: 0.55,
		text:   "X",
		onClick: func() {
			currentState = StateMenu
			w.background, _, err = ebitenutil.NewImageFromFile(BG)
			if err != nil {
				panic(err)
			}
		},
	}

	w.buttons = []*Button{buttonStart, buttonExit, buttonX}

	w.playTicker = time.NewTicker(playDuration)
	w.pauseTicker = time.NewTicker(pauseDuration)
	w.player = InitAudio()
	go w.SetSchedule()
}

func (w *Window) Update() error {
	switch currentState {
	case StateMenu:
		w.buttons[0].Update()
		w.buttons[1].Update()
	case StateGame:
		w.buttons[2].Update()
		w.gameplay.Update()
	case StateScore:
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
		w.gameplay.Draw(screen)

	case StateScore:
		data, _ := LoadGameData("data.json")
		opBG := AdjustSize(w.background, 2, 2)
		screen.DrawImage(w.scoreboard, opBG)

		w.buttons[2].Draw(screen)
		CreateRect((WIDTH-172*3)/2, HEIGHT/2-100, 3, 5, screen)
		if data != nil {
			DisplayText(WIDTH/2-175, HEIGHT/2-25, 32, fmt.Sprintf("Score: %d", data.Count), screen, color.White)
			DisplayText(WIDTH/2-175, HEIGHT/2+25, 32, fmt.Sprintf("Player: %s", data.Player), screen, color.White)
		}

	}

}

func (w *Window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return WIDTH, HEIGHT
}

func (w *Window) PlayAudio() {
	if !w.gameplay.isMusic {
		w.player.Play()
		w.gameplay.isMusic = true
		w.pauseTicker.Reset(pauseDuration)
		w.playTicker.Stop()
	}
}
func (w *Window) PauseAudio() {
	if w.gameplay.isMusic {
		w.player.Pause()
		w.gameplay.isMusic = false
		w.playTicker.Reset(playDuration)
		w.pauseTicker.Stop()
	}
}
func (w *Window) SetSchedule() {
	for {
		select {
		case <-w.pauseTicker.C:
			w.PauseAudio()
		case <-w.playTicker.C:
			w.PlayAudio()
		}
	}
}
