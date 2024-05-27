package main

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Fruit master")
	ebiten.SetTPS(TPS)

	window := Window{}
	window.Init()

	if err := ebiten.RunGame(&window); err != nil {
		log.Fatal(err)
	}
}
