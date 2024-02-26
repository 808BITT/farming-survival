package game

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	X      int
	Y      int
	Width  int
	Height int
	Img    *ebiten.Image
}

func NewTile(x, y, w, h int, img *ebiten.Image) *Tile {
	return &Tile{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
		Img:    img,
	}
}

func (t *Tile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.X), float64(t.Y))
	screen.DrawImage(t.Img, op)
}

func (t *Tile) Contains(x, y int) bool {
	return x >= t.X && x <= t.X+t.Width && y >= t.Y && y <= t.Y+t.Height
}

func (t *Tile) Visible(screenX, screenY, screenWidth, screenHeight int) bool {
	// determines if any part of the tile is visible on the screen
	return t.X < screenX+screenWidth && t.X+t.Width > screenX && t.Y < screenY+screenHeight && t.Y+t.Height > screenY
}
