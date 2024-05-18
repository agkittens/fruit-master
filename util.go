package main

import (
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func AdjustSize(img *ebiten.Image, divX int, divY int) *ebiten.DrawImageOptions {
	size := img.Bounds().Size()
	posX := (WIDTH - size.X) / divX
	posY := (HEIGHT - size.Y) / divY
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(posX), float64(posY))
	return op
}

func ChangePos(posX, posY int) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(posX), float64(posY))
	return op
}

func LoadImgs() []*ebiten.Image {
	files, _ := os.ReadDir(PATH)
	var images []*ebiten.Image
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(PATH, file.Name())
			img, _, _ := ebitenutil.NewImageFromFile(filePath)
			images = append(images, img)
		}
	}
	return images
}
