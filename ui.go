package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

	textX := b.x + (b.width-12*len(b.text))/2
	textY := b.y + (b.height-30)/2
	DisplayText(textX, textY, 24, b.text, screen)
}

type Particles struct {
}

func DisplayText(x, y, size int, msg string, screen *ebiten.Image) {
	mplusFaceSource, _ := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(color.White)

	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(size),
	}, op)
}
