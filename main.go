package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Fruit master")
	ebiten.SetTPS(TPS)

	window := Game{}

	if err := ebiten.RunGame(&window); err != nil {
		log.Fatal(err)
	}
}
