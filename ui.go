package main

import (
	"bytes"
	"image/color"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	x, y, width, height int
	scaleX, scaleY      float64
	text                string
	onClick             func()
}

func (b *Button) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		posX, posY := ebiten.CursorPosition()

		if (posX >= b.x && posX <= b.x+b.width) && (posY >= b.y && posY <= b.y+b.height) {
			b.onClick()
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image) {
	img, _, _ := ebitenutil.NewImageFromFile(BUTTON)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(b.scaleX, b.scaleY)
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(img, op)
	// col := color.RGBA{R: R, G: G, B: B, A: 255}
	// vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.width), float32(b.height), col, true)

	textX := b.x + (b.width-12*len(b.text))/2
	textY := b.y + (b.height-30)/2
	DisplayText(textX, textY, 24, b.text, screen, color.White)
}

type Particles struct {
	x, y      float32
	alpha     float32
	fadeSpeed float32
	active    bool
}

func (p *Particles) Fade() {
	if p.active {
		p.alpha -= p.fadeSpeed
		if p.alpha <= 0 {
			p.active = false
		}
	}
}

func DisplayText(x, y, size int, msg string, screen *ebiten.Image, color color.Color) {
	mplusFaceSource, _ := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	op := &text.DrawOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(color)

	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   float64(size),
	}, op)
}
