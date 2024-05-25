package main

import (
	"encoding/json"
	"image"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nfnt/resize"
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

func LoadImgs(path string) []*ebiten.Image {
	files, _ := os.ReadDir(path)
	var images []*ebiten.Image
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(path, file.Name())
			openedFile, _ := os.Open(filePath)
			img, _, _ := image.Decode(openedFile)
			w, h := 128, 128
			if file.Name() == "T_fruit_40.png" || file.Name() == "skull.png" {
				w, h = 80, 80
			}

			ebitenImg := ResizeImg(img, uint(w), uint(h))

			images = append(images, ebitenImg)
		}
	}
	return images
}

func ResizeImg(img image.Image, w, h uint) *ebiten.Image {
	resizedImg := resize.Resize(w, h, img, resize.Lanczos3)
	ebitenImg := ebiten.NewImageFromImage(resizedImg)
	return ebitenImg
}

func SaveGameData(filename string, data *GameData) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func LoadGameData(filename string) (*GameData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data GameData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
