package main

import "time"

// window consts
const WIDTH, HEIGHT = 1000, 750
const TPS = 60
const BG = "assets/bg.png"
const GAME = "assets/bg2.png"
const SCOREBOARD = "assets/bg4.png"
const TITLE = "assets/title.png"
const BUTTON = "assets/button.png"

// button color
const R, G, B = 33, 82, 117

// music
const sample = 48000
const playDuration = 3 * time.Second
const pauseDuration = 7 * time.Second

// button states
const (
	StateMenu = iota
	StateGame
	StateExit
	StateScore
)

// fruits
const PATH_F = "assets/fruitsv2/"
const PATH_B = "assets/bomb/"
