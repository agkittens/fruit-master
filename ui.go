package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type Button struct {
	x, y, width, height int
	text                string
	onClick             func()
}

func (b *Button) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		posX, posY := ebiten.CursorPosition()

		if (posX >= b.x && posX <= b.x+b.width) && (posY >= b.y && posY <= b.y+b.height) {
			log.Println("Button clicked!")
			b.onClick()
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	color := color.RGBA{R: R, G: G, B: B, A: 255}
	ebitenutil.DrawRect(screen, float64(b.x), float64(b.y), float64(b.width), float64(b.height), color)

	textWidth := text.BoundString(basicfont.Face7x13, b.text).Dx()
	textHeight := text.BoundString(basicfont.Face7x13, b.text).Dy()

	textX := b.x + (b.width-textWidth)/2
	textY := b.y + (b.height-textHeight)/2

	ebitenutil.DebugPrintAt(screen, b.text, textX, textY)
}
