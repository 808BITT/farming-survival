package assets

import (
	"embed"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed */*.png
var assets embed.FS

func LoadImage(path string) *ebiten.Image {
	f, err := assets.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf("Failed to decode image file: %v", err)
	}

	image := ebiten.NewImageFromImage(img)
	return image
}
